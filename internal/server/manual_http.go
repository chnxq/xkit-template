package server

import (
	"github.com/chnxq/xkitpkg/app"
	httptransport "github.com/chnxq/xkitpkg/transport/http"
)

func RegisterManualHTTPServices(srv *httptransport.Server, appCtx *app.AppCtx) {
	_, _ = srv, appCtx
}
