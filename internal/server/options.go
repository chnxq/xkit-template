package server

import (
	"github.com/chnxq/xkitpkg/app"
	"github.com/chnxq/xkitpkg/middleware"
)

func HTTPMiddlewares(appCtx *app.AppCtx) []middleware.Middleware {
	_ = appCtx
	return nil
}

func GRPCMiddlewares(appCtx *app.AppCtx) []middleware.Middleware {
	_ = appCtx
	return nil
}
