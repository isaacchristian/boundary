module github.com/hashicorp/boundary

go 1.21

replace github.com/hashicorp/boundary/api => ./api

replace github.com/hashicorp/boundary/sdk => ./sdk

//replace github.com/hashicorp/nodeenrollment => ../nodeenrollment

require (
	github.com/fatih/color v1.15.0
	github.com/fatih/structs v1.1.0
	github.com/favadi/protoc-go-inject-tag v1.4.0
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/golang-migrate/migrate/v4 v4.16.2
	github.com/golang-sql/civil v0.0.0-20190719163853-cb61b32ac6fe
	github.com/google/go-cmp v0.5.9
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0
	github.com/hashicorp/boundary/api v0.0.39
	github.com/hashicorp/boundary/sdk v0.0.35
	github.com/hashicorp/cap v0.3.3
	github.com/hashicorp/dawdle v0.4.0
	github.com/hashicorp/eventlogger v0.2.5
	github.com/hashicorp/eventlogger/filters/encrypt v0.1.8-0.20230428153751-cca445805f24
	github.com/hashicorp/go-bexpr v0.1.12
	github.com/hashicorp/go-cleanhttp v0.5.2
	github.com/hashicorp/go-hclog v1.5.0
	github.com/hashicorp/go-kms-wrapping/v2 v2.0.13
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/go-retryablehttp v0.7.2
	github.com/hashicorp/go-rootcerts v1.0.2
	github.com/hashicorp/go-secure-stdlib/base62 v0.1.2
	github.com/hashicorp/go-secure-stdlib/configutil/v2 v2.0.10
	github.com/hashicorp/go-secure-stdlib/gatedwriter v0.1.1
	github.com/hashicorp/go-secure-stdlib/kv-builder v0.1.2
	github.com/hashicorp/go-secure-stdlib/listenerutil v0.1.7
	github.com/hashicorp/go-secure-stdlib/mlock v0.1.3
	github.com/hashicorp/go-secure-stdlib/parseutil v0.1.7
	github.com/hashicorp/go-secure-stdlib/password v0.1.3
	github.com/hashicorp/go-secure-stdlib/pluginutil/v2 v2.0.4
	github.com/hashicorp/go-secure-stdlib/reloadutil v0.1.1
	github.com/hashicorp/go-secure-stdlib/strutil v0.1.2
	github.com/hashicorp/go-sockaddr v1.0.5
	github.com/hashicorp/go-uuid v1.0.3
	github.com/hashicorp/hcl v1.0.0
	github.com/hashicorp/vault/api v1.9.1
	github.com/iancoleman/strcase v0.2.0
	github.com/jackc/pgconn v1.14.1
	github.com/jackc/pgx/v4 v4.18.1 // indirect
	github.com/jefferai/keyring v1.1.7-0.20220316160357-58a74bb55891
	github.com/kr/pretty v0.3.1
	github.com/kr/text v0.2.0
	github.com/mattn/go-colorable v0.1.13
	github.com/mitchellh/cli v1.1.5
	github.com/mitchellh/copystructure v1.2.0
	github.com/mitchellh/go-wordwrap v1.0.1
	github.com/mitchellh/mapstructure v1.5.0
	github.com/mitchellh/pointerstructure v1.2.1
	github.com/mr-tron/base58 v1.2.0
	github.com/oligot/go-mod-upgrade v0.9.1
	github.com/ory/dockertest/v3 v3.9.1
	github.com/pires/go-proxyproto v0.7.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/posener/complete v1.2.3
	github.com/prometheus/client_golang v1.12.1
	github.com/ryanuber/go-glob v1.0.0
	github.com/stretchr/testify v1.8.4
	github.com/zalando/go-keyring v0.2.3
	go.uber.org/atomic v1.11.0
	golang.org/x/crypto v0.14.0
	golang.org/x/sync v0.2.0
	golang.org/x/sys v0.13.0
	golang.org/x/term v0.13.0
	golang.org/x/tools v0.9.1
	google.golang.org/genproto v0.0.0-20230526203410-71b5a4ffd15e
	google.golang.org/grpc v1.55.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.31.0
	gorm.io/driver/postgres v1.5.2
	gorm.io/gorm v1.25.4 // indirect
	mvdan.cc/gofumpt v0.4.0
	nhooyr.io/websocket v1.8.7
)

require github.com/hashicorp/go-dbw v0.1.0

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/cenkalti/backoff/v4 v4.2.0
	github.com/creack/pty v1.1.18
	github.com/glebarez/go-sqlite v1.21.1
	github.com/glebarez/sqlite v1.8.0
	github.com/golang/protobuf v1.5.3
	github.com/hashicorp/cap/ldap v0.0.0-20230918191608-899e3f8e8cd6
	github.com/hashicorp/dbassert v0.0.0-20230622135851-cc4c0f18f4a7
	github.com/hashicorp/go-kms-wrapping/extras/kms/v2 v2.0.0-20230902160534-7a966e6313fd
	github.com/hashicorp/go-version v1.6.0
	github.com/hashicorp/nodeenrollment v0.2.7
	github.com/jackc/pgx/v5 v5.4.3
	github.com/jimlambrt/gldap v0.1.7
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/mikesmitty/edkey v0.0.0-20170222072505-3356ea4e686a
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/go-ps v1.0.0
	github.com/sevlyar/go-daemon v0.1.6
	golang.org/x/exp v0.0.0-20230811145659-89c5cff77bcb
	golang.org/x/net v0.15.0
	google.golang.org/genproto/googleapis/api v0.0.0-20230530153820-e85fd2cbaebc
	gopkg.in/square/go-jose.v2 v2.5.1
)

require (
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
)

require (
	github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4 // indirect
	github.com/AlecAivazis/survey/v2 v2.3.2 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161 // indirect
	github.com/Azure/go-ntlmssp v0.0.0-20221128193559-754e69321358 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.2.0 // indirect
	github.com/Masterminds/sprig/v3 v3.2.3 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/alessio/shellescape v1.4.1 // indirect
	github.com/apex/log v1.9.0 // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bgentry/speakeasy v0.1.0 // indirect
	github.com/briandowns/spinner v1.16.0 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cenkalti/backoff/v3 v3.0.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/containerd/continuity v0.3.0 // indirect
	github.com/coreos/go-oidc/v3 v3.5.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/danieljoos/wincred v1.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/cli v23.0.1+incompatible // indirect
	github.com/docker/docker v23.0.3+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/dvsekhvalnov/jose2go v1.5.0 // indirect
	github.com/go-asn1-ber/asn1-ber v1.5.5 // indirect
	github.com/go-jose/go-jose/v3 v3.0.0 // indirect
	github.com/go-ldap/ldap/v3 v3.4.6 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v1.1.0 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-kms-wrapping/plugin/v2 v2.0.4 // indirect
	github.com/hashicorp/go-plugin v1.4.9 // indirect
	github.com/hashicorp/go-secure-stdlib/tlsutil v0.1.2 // indirect
	github.com/hashicorp/mql v0.0.0-20230814234610-74af33a21d7f
	github.com/hashicorp/vault/sdk v0.3.0 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/huandu/xstrings v1.4.0 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.2 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.14.0 // indirect
	github.com/jefferai/go-libsecret v0.0.0-20210525195240-b53481abef97 // indirect
	github.com/jefferai/isbadcipher v0.0.0-20190226160619-51d2077c035f // indirect
	github.com/jinzhu/gorm v1.9.12 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/klauspost/compress v1.15.11 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/term v0.5.0 // indirect
	github.com/mtibben/percent v0.2.1 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/opencontainers/runc v1.1.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/sethvargo/go-diceware v0.3.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/sirupsen/logrus v1.9.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/urfave/cli/v2 v2.3.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/xo/dburl v0.16.0 // indirect
	go.uber.org/goleak v1.1.12 // indirect
	golang.org/x/mod v0.11.0 // indirect
	golang.org/x/oauth2 v0.8.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/sqlite v1.5.3 // indirect
	modernc.org/libc v1.22.3 // indirect
	modernc.org/mathutil v1.5.0 // indirect
	modernc.org/memory v1.5.0 // indirect
	modernc.org/sqlite v1.21.1 // indirect
)

replace nhooyr.io/websocket => github.com/tmessi/websocket v0.0.0-20231006132118-662e2f7878ac
