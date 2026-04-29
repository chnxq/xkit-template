package server

import (
	"fmt"

	"github.com/gorilla/handlers"

	"github.com/chnxq/xkitpkg/app"
	conf "github.com/chnxq/xkitpkg/conf/v1"
	httptransport "github.com/chnxq/xkitpkg/transport/http"
	httppprof "github.com/chnxq/xkitpkg/transport/http/pprof"
)

func HTTPServerOptions(appCtx *app.AppCtx) ([]httptransport.ServerOption, error) {
	cfg := restConfig(appCtx)
	opts := []httptransport.ServerOption{
		httptransport.Address(":8000"),
	}

	if cfg != nil {
		if cfg.GetNetwork() != "" {
			opts = append(opts, httptransport.Network(cfg.GetNetwork()))
		}
		if cfg.GetAddr() != "" {
			opts = append(opts, httptransport.Address(cfg.GetAddr()))
		}
		if cfg.GetTimeout() != nil {
			opts = append(opts, httptransport.Timeout(cfg.GetTimeout().AsDuration()))
		}
		if cfg.GetTls() != nil {
			tlsConfig, err := loadServerTLSConfig(cfg.GetTls())
			if err != nil {
				return nil, fmt.Errorf("load rest tls config: %w", err)
			}
			if tlsConfig != nil {
				opts = append(opts, httptransport.TLSConfig(tlsConfig))
			}
		}
	}

	if middlewares := append(commonServerMiddlewares(appCtx, cfg.GetMiddleware()), HTTPMiddlewares(appCtx)...); len(middlewares) > 0 {
		opts = append(opts, httptransport.Middleware(middlewares...))
	}
	if filters := restConfigFilters(cfg); len(filters) > 0 {
		opts = append(opts, httptransport.Filter(filters...))
	}
	return opts, nil
}

func RegisterConfiguredHTTPHandlers(srv *httptransport.Server, appCtx *app.AppCtx) {
	if srv == nil {
		return
	}
	if cfg := restConfig(appCtx); cfg == nil || !cfg.GetEnablePprof() {
		return
	}

	handler := httppprof.NewHandler()
	srv.Handle("/debug/pprof", handler)
	srv.HandlePrefix("/debug/pprof/", handler)
}

func restConfigFilters(cfg *conf.Server_REST) []httptransport.FilterFunc {
	if cfg == nil || cfg.GetCors() == nil {
		return nil
	}

	cors := cfg.GetCors()
	return []httptransport.FilterFunc{
		handlers.CORS(
			handlers.AllowedHeaders(cors.GetHeaders()),
			handlers.AllowedMethods(cors.GetMethods()),
			handlers.AllowedOrigins(cors.GetOrigins()),
		),
	}
}
