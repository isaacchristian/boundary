// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package ratelimit

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"regexp"

	"github.com/hashicorp/boundary/globals"
	"github.com/hashicorp/boundary/internal/event"
	"github.com/hashicorp/boundary/internal/types/action"
	"github.com/hashicorp/boundary/internal/types/resource"
	"github.com/hashicorp/go-rate"
)

var (
	errUnknownResource   = &extractResourceActionErr{http.StatusNotFound, "unknown resource"}
	errUnknownAction     = &extractResourceActionErr{http.StatusBadRequest, "unknown action"}
	errUnsupportedAction = &extractResourceActionErr{http.StatusMethodNotAllowed, "invalid action"}
)

type extractResourceActionErr struct {
	statusCode int
	msg        string
}

func (e *extractResourceActionErr) Error() string {
	return e.msg
}

var pathRegex = regexp.MustCompile(`/v1/(?P<resource>[\w-]+)((/(?P<id>[\w/]+))?(:(?P<action>[\w-:]+)?)?)?`)

func extractResourceAction(path, method string) (res, act string, err error) {
	var id string

	var r resource.Type
	var ok bool
	var actionSet action.ActionSet

	// TODO: replace regex with lexer
	match := pathRegex.FindStringSubmatch(path)
	for i, name := range pathRegex.SubexpNames() {
		switch name {
		case "resource":
			res = match[i]
			r, ok = resource.FromPlural(res)
			if !ok {
				res = resource.Unknown.String()
				err = errUnknownResource
				return
			}
			res = r.String()
		case "action":
			act = match[i]
			if act != "" {
				actionSet, err = action.ActionSetForResource(r)
				if err != nil {
					act = action.Unknown.String()
					err = errUnknownAction
					return
				}
				at, ok := action.Map[act]
				if !ok {
					act = action.Unknown.String()
					err = errUnknownAction
					return
				}
				if !actionSet.HasAction(at) {
					err = errUnsupportedAction
					return
				}
			}
		case "id":
			id = match[i]
		}
	}

	switch act {
	case "":
		switch id {
		case "":
			switch method {
			case http.MethodGet:
				act = action.List.String()
			case http.MethodPost:
				act = action.Create.String()
			default:
				act = action.Unknown.String()
				err = errUnsupportedAction
				return
			}
		default:
			switch method {
			case http.MethodGet:
				act = action.Read.String()
			case http.MethodDelete:
				act = action.Delete.String()
			case http.MethodPatch:
				act = action.Update.String()
			default:
				act = action.Unknown.String()
				err = errUnsupportedAction
				return
			}
		}
	}
	return
}

// Handler is an http middleware handler that checks if a request is allowed
// using the provided rate limiter. If the request is allowed, the next handler
// is called. Otherwise a 429 is returned with the Retry-After response header
// set to the number of seconds the client should wait to make it's next request.
func Handler(ctx context.Context, l *rate.Limiter, next http.Handler) http.Handler {
	const op = "ratelimit.Handler"

	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		reqInfo, ok := event.RequestInfoFromContext(req.Context())
		if !ok || reqInfo == nil || reqInfo.ClientIp == "" {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		authtoken, ok := req.Context().Value(globals.ContextAuthTokenPublicIdKey).(string)
		if !ok {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		res, a, err := extractResourceAction(req.URL.Path, req.Method)
		if err != nil {
			if extractErr, ok := err.(*extractResourceActionErr); ok {
				rw.WriteHeader(extractErr.statusCode)
				return
			}
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		allowed, quota, err := l.Allow(res, a, reqInfo.ClientIp, authtoken)
		if err != nil {
			if errFull, ok := err.(*rate.ErrLimiterFull); ok {
				rw.Header().Add("Retry-After", fmt.Sprintf("%.0f", math.Ceil(errFull.RetryIn.Seconds())))
				rw.WriteHeader(http.StatusServiceUnavailable)
				return
			}
			// The only other error here should be rate.ErrLimitNotFound, which
			// shouldn't be possible given how we initialize the limiter and
			// the checks done by extractResourceAction.
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !allowed {
			rw.Header().Add("Retry-After", fmt.Sprintf("%.0f", quota.ResetsIn().Seconds()))
			rw.WriteHeader(http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(rw, req)
	})
}