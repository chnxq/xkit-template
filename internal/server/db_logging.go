package server

import "github.com/chnxq/xkitpkg/middleware"

type DatabaseLoggingData interface {
	DatabaseLoggingMiddleware() middleware.Middleware
}

func databaseLoggingMiddleware(data GeneratedData) middleware.Middleware {
	if provider, ok := data.(DatabaseLoggingData); ok {
		if mw := provider.DatabaseLoggingMiddleware(); mw != nil {
			return mw
		}
	}
	return func(next middleware.Handler) middleware.Handler {
		return next
	}
}
