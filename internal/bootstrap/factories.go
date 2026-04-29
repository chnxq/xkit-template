package bootstrap

// Preload factory registrations used by config, logger, and registry creation.
import (
	_ "github.com/chnxq/xkitpkg/config/consul"
	_ "github.com/chnxq/xkitpkg/config/etcd"
	_ "github.com/chnxq/xkitpkg/logger/fluentd"
	_ "github.com/chnxq/xkitpkg/logger/zap"
	_ "github.com/chnxq/xkitpkg/registry/consul"
	_ "github.com/chnxq/xkitpkg/registry/etcd"
)
