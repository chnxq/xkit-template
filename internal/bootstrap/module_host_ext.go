package bootstrap

import (
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
	return modulehost.HostServices{
		Context: appCtx.AppContext(),
		Loggers: appLoggerFactory{appCtx: appCtx},
	}
}
