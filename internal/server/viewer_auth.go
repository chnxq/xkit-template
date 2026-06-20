package server

import "github.com/chnxq/xkitpkg/middleware"

func authViewerMiddleware(data GeneratedData) middleware.Middleware {
	_ = data
	return nil
}
