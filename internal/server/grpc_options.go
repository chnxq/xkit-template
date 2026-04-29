package server

import (
	"fmt"

	"github.com/chnxq/xkitpkg/app"
	grpctransport "github.com/chnxq/xkitpkg/transport/grpc"
)

func GRPCServerOptions(appCtx *app.AppCtx) ([]grpctransport.ServerOption, error) {
	cfg := grpcConfig(appCtx)
	opts := []grpctransport.ServerOption{
		grpctransport.Address(":9000"),
	}

	if cfg != nil {
		if cfg.GetNetwork() != "" {
			opts = append(opts, grpctransport.Network(cfg.GetNetwork()))
		}
		if cfg.GetAddr() != "" {
			opts = append(opts, grpctransport.Address(cfg.GetAddr()))
		}
		if cfg.GetTimeout() != nil {
			opts = append(opts, grpctransport.Timeout(cfg.GetTimeout().AsDuration()))
		}
		if cfg.GetTls() != nil {
			tlsConfig, err := loadServerTLSConfig(cfg.GetTls())
			if err != nil {
				return nil, fmt.Errorf("load grpc tls config: %w", err)
			}
			if tlsConfig != nil {
				opts = append(opts, grpctransport.TLSConfig(tlsConfig))
			}
		}
	}

	if middlewares := append(commonServerMiddlewares(appCtx, cfg.GetMiddleware()), GRPCMiddlewares(appCtx)...); len(middlewares) > 0 {
		opts = append(opts, grpctransport.Middleware(middlewares...))
	}
	return opts, nil
}
