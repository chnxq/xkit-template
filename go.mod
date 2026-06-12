module xkit-template-v01

replace github.com/armon/go-metrics => github.com/hashicorp/go-metrics v0.4.1

go 1.26.0

require (
	github.com/chnxq/x-crud/viewer v0.0.0-20260411151944-a61448f9f7bc
	github.com/chnxq/x-swagger v0.0.0-20260529105209-02745c8a5170
	github.com/chnxq/x-utils v0.0.0-20260612100514-4160a415201a
	github.com/chnxq/xkitmod v0.0.0-20260529105211-b1cd4c65f020
	github.com/chnxq/xkitmod/algs v0.0.0-20260529105211-b1cd4c65f020
	github.com/chnxq/xkitmod/log v0.0.0-20260529105211-b1cd4c65f020
	github.com/chnxq/xkitpkg/app v0.0.0-20260612150007-4a0b93a9efd2
	github.com/chnxq/xkitpkg/conf v0.0.0-20260612150007-4a0b93a9efd2
	github.com/chnxq/xkitpkg/config v0.0.0-20260612150007-4a0b93a9efd2
	github.com/chnxq/xkitpkg/config/consul v0.0.0-20260612150007-4a0b93a9efd2
	github.com/chnxq/xkitpkg/config/etcd v0.0.0-20260612150007-4a0b93a9efd2
	github.com/chnxq/xkitpkg/logger v0.0.0-20260612150007-4a0b93a9efd2
	github.com/chnxq/xkitpkg/logger/fluentd v0.0.0-20260612150007-4a0b93a9efd2
	github.com/chnxq/xkitpkg/logger/zap v0.0.0-20260612150007-4a0b93a9efd2
	github.com/chnxq/xkitpkg/middleware v0.0.0-20260612150007-4a0b93a9efd2
	github.com/chnxq/xkitpkg/registry v0.0.0-20260612150007-4a0b93a9efd2
	github.com/chnxq/xkitpkg/registry/consul v0.0.0-20260612150007-4a0b93a9efd2
	github.com/chnxq/xkitpkg/registry/etcd v0.0.0-20260612150007-4a0b93a9efd2
	github.com/chnxq/xkitpkg/tracer v0.0.0-20260612150007-4a0b93a9efd2
	github.com/chnxq/xkitpkg/transport v0.0.0-20260612150007-4a0b93a9efd2
	github.com/gorilla/handlers v1.5.2
	google.golang.org/grpc v1.81.1
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.11-20260415201107-50325440f8f2.1 // indirect
	buf.build/go/protovalidate v1.2.0 // indirect
	cel.dev/expr v0.25.2 // indirect
	dario.cat/mergo v1.0.2 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/armon/go-metrics v0.5.4 // indirect
	github.com/cenkalti/backoff/v5 v5.0.3 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/chnxq/xkitmod/config v0.0.0-20260529105211-b1cd4c65f020 // indirect
	github.com/chnxq/xkitmod/selector v0.0.0-20260529105211-b1cd4c65f020 // indirect
	github.com/chnxq/xkitpkg v0.0.0-20260612150007-4a0b93a9efd2 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.7.0 // indirect
	github.com/fatih/color v1.19.0 // indirect
	github.com/felixge/httpsnoop v1.1.0 // indirect
	github.com/fluent/fluent-logger-golang v1.10.1 // indirect
	github.com/fsnotify/fsnotify v1.10.1 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-playground/form/v4 v4.3.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.5.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/cel-go v0.28.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.29.0 // indirect
	github.com/hashicorp/consul/api v1.34.3 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-metrics v0.5.4 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/serf v0.10.2 // indirect
	github.com/hibiken/asynq v0.26.0 // indirect
	github.com/lufia/plan9stats v0.0.0-20260330125221-c963978e514e // indirect
	github.com/mattn/go-colorable v0.1.15 // indirect
	github.com/mattn/go-isatty v0.0.22 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/openzipkin/zipkin-go v0.4.3 // indirect
	github.com/philhofer/fwd v1.2.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20240221224432-82ca36839d55 // indirect
	github.com/redis/go-redis/v9 v9.20.1 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/shirou/gopsutil/v3 v3.24.5 // indirect
	github.com/shoenig/go-m1cpu v0.2.1 // indirect
	github.com/spf13/cast v1.10.0 // indirect
	github.com/swaggest/swgui v1.8.8 // indirect
	github.com/tinylib/msgp v1.6.4 // indirect
	github.com/tklauser/go-sysconf v0.4.0 // indirect
	github.com/tklauser/numcpus v0.12.0 // indirect
	github.com/vearutop/statigz v1.5.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.etcd.io/etcd/api/v3 v3.6.12 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.6.12 // indirect
	go.etcd.io/etcd/client/v3 v3.6.12 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/contrib/bridges/otelzap v0.19.0 // indirect
	go.opentelemetry.io/otel v1.44.0 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.17.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc v0.20.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.44.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.44.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.44.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.44.0 // indirect
	go.opentelemetry.io/otel/exporters/zipkin v1.44.0 // indirect
	go.opentelemetry.io/otel/log v0.20.0 // indirect
	go.opentelemetry.io/otel/metric v1.44.0 // indirect
	go.opentelemetry.io/otel/sdk v1.44.0 // indirect
	go.opentelemetry.io/otel/sdk/log v0.20.0 // indirect
	go.opentelemetry.io/otel/trace v1.44.0 // indirect
	go.opentelemetry.io/proto/otlp v1.10.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.28.0 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/exp v0.0.0-20260611194520-c48552f49976 // indirect
	golang.org/x/net v0.56.0 // indirect
	golang.org/x/sync v0.21.0 // indirect
	golang.org/x/sys v0.46.0 // indirect
	golang.org/x/text v0.38.0 // indirect
	golang.org/x/time v0.15.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20260610212136-7ab31c22f7ad // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260610212136-7ab31c22f7ad // indirect
	google.golang.org/protobuf v1.36.11 // indirect
	gopkg.in/cenkalti/backoff.v1 v1.1.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
