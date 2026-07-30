package bootstrap

import (
	"context"
	"time"

	"xkit-template-v01/internal/server"
	_ "xkit-template-v01/modules"
	modulehost "xkit-template-v01/shared/modulehost"

	"github.com/chnxq/xkitpkg/app"
	httptransport "github.com/chnxq/xkitpkg/transport/http"
	"google.golang.org/grpc"
)

const moduleShutdownTimeout = 30 * time.Second

type hostModuleRuntime struct {
	runtime modulehost.Runtime
}

func (r hostModuleRuntime) RegisterHTTP(srv *httptransport.Server) {
	if r.runtime != nil {
		r.runtime.RegisterHTTP(srv)
	}
}

func (r hostModuleRuntime) RegisterGRPC(registrar grpc.ServiceRegistrar) {
	if r.runtime != nil {
		r.runtime.RegisterGRPC(registrar)
	}
}

func loadHostModules(appCtx *app.AppCtx, cleanup *CleanupStack) ([]hostModuleRuntime, error) {
	if appCtx == nil {
		return nil, nil
	}

	hostServices := BuildModuleHostServices(appCtx)
	definitions := modulehost.Definitions()
	runtimes := make([]hostModuleRuntime, 0, len(definitions))
	for _, definition := range definitions {
		if definition == nil {
			continue
		}
		runtime, err := definition.BuildRuntime(hostServices)
		if err != nil {
			return nil, err
		}
		if runtime == nil {
			continue
		}
		runtimes = append(runtimes, hostModuleRuntime{runtime: runtime})
		cleanup.AddContext(func(parent context.Context) error {
			ctx, cancel := context.WithTimeout(parent, moduleShutdownTimeout)
			defer cancel()
			err := runtime.Close(ctx)
			if err != nil && appCtx != nil {
				appCtx.NewLoggerHelper("module/runtime").Errorf(
					"close module runtime failed: module=%s error=%v",
					runtime.Descriptor().Code,
					err,
				)
			}
			return err
		})
	}
	return runtimes, nil
}

func SyncHostModuleResources(ctx context.Context, syncer modulehost.ResourceSyncer) error {
	if syncer == nil {
		return nil
	}
	for _, definition := range modulehost.Definitions() {
		if definition == nil {
			continue
		}
		menus := definition.Resources().Menus
		if len(menus) == 0 {
			continue
		}
		if err := syncer.UpsertMenus(ctx, menus); err != nil {
			return err
		}
	}
	return nil
}

func hostModuleHTTPRegistrars(runtimes []hostModuleRuntime) []server.ModuleHTTPRegistrar {
	if len(runtimes) == 0 {
		return nil
	}
	items := make([]server.ModuleHTTPRegistrar, 0, len(runtimes))
	for _, runtime := range runtimes {
		items = append(items, runtime)
	}
	return items
}

func hostModuleGRPCRegistrars(runtimes []hostModuleRuntime) []server.ModuleGRPCRegistrar {
	if len(runtimes) == 0 {
		return nil
	}
	items := make([]server.ModuleGRPCRegistrar, 0, len(runtimes))
	for _, runtime := range runtimes {
		items = append(items, runtime)
	}
	return items
}
