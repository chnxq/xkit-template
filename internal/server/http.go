package server

import (
	"fmt"

	"github.com/chnxq/xkitpkg/app"
	httptransport "github.com/chnxq/xkitpkg/transport/http"
)

func NewHTTPServer(appCtx *app.AppCtx, services GeneratedHTTPServices) (*httptransport.Server, error) {
	opts, err := HTTPServerOptions(appCtx)
	if err != nil {
		return nil, fmt.Errorf("http server options: %w", err)
	}
	srv := httptransport.NewServer(opts...)
	RegisterGeneratedHTTPServices(srv, services)
	//RegisterManualHTTPServices(srv, appCtx)
	RegisterConfiguredHTTPHandlers(srv, appCtx)
	return srv, nil
}
