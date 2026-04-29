package server

import (
	"github.com/chnxq/xkitpkg/app"
	conf "github.com/chnxq/xkitpkg/conf/v1"
	ssetransport "github.com/chnxq/xkitpkg/transport/sse"
)

func NewSSEServer(appCtx *app.AppCtx) (*ssetransport.Server, error) {
	if appCtx == nil {
		return nil, nil
	}
	cfg := appCtx.GetConfig()
	if cfg == nil || cfg.Server == nil || cfg.Server.Sse == nil {
		return nil, nil
	}

	srv, err := newConfiguredSSEServer(appCtx, cfg.Server.Sse)
	if err != nil {
		return nil, err
	}
	if srv == nil {
		return nil, nil
	}
	//if err := RegisterManualSSEHandlers(srv, appCtx); err != nil {
	//	return nil, err
	//}
	return srv, nil
}

func newConfiguredSSEServer(appCtx *app.AppCtx, cfg *conf.Server_SSE) (*ssetransport.Server, error) {
	if cfg == nil {
		return nil, nil
	}
	opts, err := SSEServerOptions(appCtx, cfg)
	if err != nil {
		return nil, err
	}
	return ssetransport.NewServer(opts...), nil
}
