package bootstrap

import (
	"time"

	modulehost "xkit-template-v01/shared/modulehost"

	"github.com/chnxq/xkitpkg/app"
)

type appLoggerFactory struct {
	appCtx *app.AppCtx
}

func (f appLoggerFactory) NewLogger(name string) modulehost.Logger {
	return f.appCtx.NewLoggerHelper(name)
}

// BuildModuleHostServices is the host-owned composition root for module ports.
func BuildModuleHostServices(appCtx *app.AppCtx) modulehost.HostServices {
	if appCtx == nil {
		return modulehost.HostServices{}
	}
	services := modulehost.HostServices{
		Context: appCtx.AppContext(),
		Loggers: appLoggerFactory{appCtx: appCtx},
	}
	if cfg := appCtx.GetConfig(); cfg != nil && cfg.GetData() != nil && cfg.GetData().GetDatabase() != nil {
		database := cfg.GetData().GetDatabase()
		services.PrimaryDataSource = modulehost.DeploymentSpec{
			DeploymentVersion:    1,
			Driver:               database.GetDriver(),
			Source:               database.GetSource(),
			ConfigurationVersion: 1,
			Pool: modulehost.PoolPolicy{
				MaxIdleConnections: positiveUint32(database.GetMaxIdleConnections()),
				MaxOpenConnections: positiveUint32(database.GetMaxOpenConnections()),
				ConnectionMaxLife:  connectionMaxLifetime(database.GetConnectionMaxLifetime()),
			},
		}
	}
	return services
}

func positiveUint32(value int32) uint32 {
	if value <= 0 {
		return 0
	}
	return uint32(value)
}

func connectionMaxLifetime(value interface{ AsDuration() time.Duration }) time.Duration {
	if value == nil {
		return 0
	}
	return value.AsDuration()
}
