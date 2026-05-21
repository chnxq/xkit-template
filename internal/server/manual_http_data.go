package server

import (
	"github.com/chnxq/xkitpkg/app"
	httptransport "github.com/chnxq/xkitpkg/transport/http"
)

// RegisterManualHTTPServicesWithData is the template-owned hook for manual
// HTTP handlers that need GeneratedData without changing preserved
// RegisterManualHTTPServices signatures in existing projects.
func RegisterManualHTTPServicesWithData(srv *httptransport.Server, appCtx *app.AppCtx, data GeneratedData) {
	_, _, _ = srv, appCtx, data
}
