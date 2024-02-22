module storj.io/common

go 1.20

require (
	cloud.google.com/go/profiler v0.4.0
	github.com/blang/semver v3.5.1+incompatible
	github.com/bmkessler/fastdiv v0.0.0-20190227075523-41d5178f2044
	github.com/calebcase/tmpfile v1.0.3
	github.com/flynn/noise v1.0.0
	github.com/gogo/protobuf v1.3.2
	github.com/google/pprof v0.0.0-20230602150820-91b7bce49751
	github.com/jackc/pgtype v1.14.1
	github.com/jackc/pgx/v5 v5.5.2
	github.com/jtolds/tracetagger/v2 v2.0.0-rc5
	github.com/jtolio/crawlspace v0.0.0-20231116162947-3ec5cc6b36c5
	github.com/jtolio/crawlspace/tools v0.0.0-20231116162947-3ec5cc6b36c5
	github.com/jtolio/noiseconn v0.0.0-20230111204749-d7ec1a08b0b8
	github.com/mattn/go-sqlite3 v1.14.19
	github.com/quic-go/quic-go v0.40.1
	github.com/shopspring/decimal v1.2.0
	github.com/spacemonkeygo/monkit/v3 v3.0.22
	github.com/spf13/cast v1.6.0
	github.com/spf13/cobra v1.8.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.18.2
	github.com/stretchr/testify v1.8.4
	github.com/zeebo/admission/v3 v3.0.3
	github.com/zeebo/blake3 v0.2.3
	github.com/zeebo/errs v1.3.0
	github.com/zeebo/structs v1.0.2
	go.uber.org/zap v1.26.0
	golang.org/x/crypto v0.17.0
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9
	golang.org/x/sync v0.5.0
	golang.org/x/sys v0.15.0
	gopkg.in/yaml.v2 v2.4.0
	storj.io/drpc v0.0.33
	storj.io/eventkit v0.0.0-20240222101345-204fc04c7675
	storj.io/monkit-jaeger v0.0.0-20240221095020-52b0792fa6cd
	storj.io/picobuf v0.0.3
)

require (
	cloud.google.com/go v0.110.10 // indirect
	cloud.google.com/go/compute v1.23.3 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/apache/thrift v0.16.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/klauspost/cpuid/v2 v2.0.12 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/onsi/ginkgo/v2 v2.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/quic-go/qtls-go1-20 v0.4.1 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/zeebo/float16 v0.1.0 // indirect
	github.com/zeebo/goof v0.0.0-20230907150950-e9457bc94477 // indirect
	github.com/zeebo/incenc v0.0.0-20180505221441-0d92902eec54 // indirect
	github.com/zeebo/mwc v0.0.4 // indirect
	github.com/zeebo/sudo v1.0.2 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.uber.org/mock v0.3.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/oauth2 v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.13.0 // indirect
	google.golang.org/api v0.153.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20231106174013-bbf56f31fb17 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231106174013-bbf56f31fb17 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231120223509-83a465c0220f // indirect
	google.golang.org/grpc v1.59.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
