// Code generated by "make api"; DO NOT EDIT.
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedgroups

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/scopes"
	"golang.org/x/exp/slices"
)

type ManagedGroup struct {
	Id                string                 `json:"id,omitempty"`
	Scope             *scopes.ScopeInfo      `json:"scope,omitempty"`
	Name              string                 `json:"name,omitempty"`
	Description       string                 `json:"description,omitempty"`
	CreatedTime       time.Time              `json:"created_time,omitempty"`
	UpdatedTime       time.Time              `json:"updated_time,omitempty"`
	Version           uint32                 `json:"version,omitempty"`
	Type              string                 `json:"type,omitempty"`
	AuthMethodId      string                 `json:"auth_method_id,omitempty"`
	Attributes        map[string]interface{} `json:"attributes,omitempty"`
	MemberIds         []string               `json:"member_ids,omitempty"`
	AuthorizedActions []string               `json:"authorized_actions,omitempty"`

	response *api.Response
}

type ManagedGroupReadResult struct {
	Item     *ManagedGroup
	response *api.Response
}

func (n ManagedGroupReadResult) GetItem() *ManagedGroup {
	return n.Item
}

func (n ManagedGroupReadResult) GetResponse() *api.Response {
	return n.response
}

type ManagedGroupCreateResult = ManagedGroupReadResult
type ManagedGroupUpdateResult = ManagedGroupReadResult

type ManagedGroupDeleteResult struct {
	response *api.Response
}

// GetItem will always be nil for ManagedGroupDeleteResult
func (n ManagedGroupDeleteResult) GetItem() interface{} {
	return nil
}

func (n ManagedGroupDeleteResult) GetResponse() *api.Response {
	return n.response
}

type ManagedGroupListResult struct {
	Items        []*ManagedGroup `json:"items,omitempty"`
	EstItemCount uint            `json:"est_item_count,omitempty"`
	RemovedIds   []string        `json:"removed_ids,omitempty"`
	RefreshToken string          `json:"refresh_token,omitempty"`
	ResponseType string          `json:"response_type,omitempty"`
	response     *api.Response
}

func (n ManagedGroupListResult) GetItems() []*ManagedGroup {
	return n.Items
}

func (n ManagedGroupListResult) GetEstItemCount() uint {
	return n.EstItemCount
}

func (n ManagedGroupListResult) GetRemovedIds() []string {
	return n.RemovedIds
}

func (n ManagedGroupListResult) GetRefreshToken() string {
	return n.RefreshToken
}

func (n ManagedGroupListResult) GetResponseType() string {
	return n.ResponseType
}

func (n ManagedGroupListResult) GetResponse() *api.Response {
	return n.response
}

// Client is a client for this collection
type Client struct {
	client *api.Client
}

// Creates a new client for this collection. The submitted API client is cloned;
// modifications to it after generating this client will not have effect. If you
// need to make changes to the underlying API client, use ApiClient() to access
// it.
func NewClient(c *api.Client) *Client {
	return &Client{client: c.Clone()}
}

// ApiClient returns the underlying API client
func (c *Client) ApiClient() *api.Client {
	return c.client
}

func (c *Client) Create(ctx context.Context, authMethodId string, opt ...Option) (*ManagedGroupCreateResult, error) {
	if authMethodId == "" {
		return nil, fmt.Errorf("empty authMethodId value passed into Create request")
	}

	opts, apiOpts := getOpts(opt...)

	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts.postMap["auth_method_id"] = authMethodId

	req, err := c.client.NewRequest(ctx, "POST", "managed-groups", opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Create request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Create call: %w", err)
	}

	target := new(ManagedGroupCreateResult)
	target.Item = new(ManagedGroup)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Create response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) Read(ctx context.Context, id string, opt ...Option) (*ManagedGroupReadResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into Read request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("managed-groups/%s", url.PathEscape(id)), nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Read request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Read call: %w", err)
	}

	target := new(ManagedGroupReadResult)
	target.Item = new(ManagedGroup)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Read response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) Update(ctx context.Context, id string, version uint32, opt ...Option) (*ManagedGroupUpdateResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into Update request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into Update request and automatic versioning not specified")
		}
		existingTarget, existingErr := c.Read(ctx, id, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version

	req, err := c.client.NewRequest(ctx, "PATCH", fmt.Sprintf("managed-groups/%s", url.PathEscape(id)), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Update request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Update call: %w", err)
	}

	target := new(ManagedGroupUpdateResult)
	target.Item = new(ManagedGroup)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Update response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) Delete(ctx context.Context, id string, opt ...Option) (*ManagedGroupDeleteResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into Delete request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("managed-groups/%s", url.PathEscape(id)), nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Delete request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Delete call: %w", err)
	}

	apiErr, err := resp.Decode(nil)
	if err != nil {
		return nil, fmt.Errorf("error decoding Delete response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}

	target := &ManagedGroupDeleteResult{
		response: resp,
	}
	return target, nil
}

func (c *Client) List(ctx context.Context, authMethodId string, opt ...Option) (*ManagedGroupListResult, error) {
	if authMethodId == "" {
		return nil, fmt.Errorf("empty authMethodId value passed into List request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	opts.queryMap["auth_method_id"] = authMethodId

	req, err := c.client.NewRequest(ctx, "GET", "managed-groups", nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating List request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during List call: %w", err)
	}

	target := new(ManagedGroupListResult)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, fmt.Errorf("error decoding List response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	if target.ResponseType == "complete" || target.ResponseType == "" {
		return target, nil
	}
	// If there are more results, automatically fetch the rest of the results.
	// idToIndex keeps a map from the ID of an item to its index in target.Items.
	// This is used to update updated items in-place and remove deleted items
	// from the result after pagination is done.
	idToIndex := map[string]int{}
	for i, item := range target.Items {
		idToIndex[item.Id] = i
	}
	for {
		req, err := c.client.NewRequest(ctx, "GET", "managed-groups", nil, apiOpts...)
		if err != nil {
			return nil, fmt.Errorf("error creating List request: %w", err)
		}

		opts.queryMap["refresh_token"] = target.RefreshToken
		if len(opts.queryMap) > 0 {
			q := url.Values{}
			for k, v := range opts.queryMap {
				q.Add(k, v)
			}
			req.URL.RawQuery = q.Encode()
		}

		resp, err := c.client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error performing client request during List call: %w", err)
		}

		page := new(ManagedGroupListResult)
		apiErr, err := resp.Decode(page)
		if err != nil {
			return nil, fmt.Errorf("error decoding List response: %w", err)
		}
		if apiErr != nil {
			return nil, apiErr
		}
		for _, item := range page.Items {
			if i, ok := idToIndex[item.Id]; ok {
				// Item has already been seen at index i, update in-place
				target.Items[i] = item
			} else {
				target.Items = append(target.Items, item)
				idToIndex[item.Id] = len(target.Items) - 1
			}
		}
		target.RemovedIds = append(target.RemovedIds, page.RemovedIds...)
		target.EstItemCount = page.EstItemCount
		target.RefreshToken = page.RefreshToken
		target.ResponseType = page.ResponseType
		target.response = resp
		if target.ResponseType == "complete" {
			break
		}
	}
	for _, removedId := range target.RemovedIds {
		if i, ok := idToIndex[removedId]; ok {
			// Remove the item at index i without preserving order
			// https://github.com/golang/go/wiki/SliceTricks#delete-without-preserving-order
			target.Items[i] = target.Items[len(target.Items)-1]
			target.Items = target.Items[:len(target.Items)-1]
			// Update the index of the last element
			idToIndex[target.Items[i].Id] = i
		}
	}
	// Finally, sort the results again since in-place updates and deletes
	// may have shuffled items.
	slices.SortFunc(target.Items, func(i, j *ManagedGroup) int {
		return i.UpdatedTime.Compare(j.UpdatedTime)
	})
	return target, nil
}
