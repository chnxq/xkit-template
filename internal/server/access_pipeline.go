package server

import "github.com/chnxq/xkitpkg/middleware"

type AccessPipeline struct {
	Authentication   middleware.Middleware
	TenantScope      middleware.Middleware
	AccessContext    middleware.Middleware
	Authorization    middleware.Middleware
	ModuleDeployment middleware.Middleware
}

func BuildAccessPipeline(data GeneratedData) AccessPipeline {
	_ = data
	return AccessPipeline{}
}

func appendAccessPipeline(items []middleware.Middleware, pipeline AccessPipeline) []middleware.Middleware {
	for _, item := range []middleware.Middleware{
		pipeline.Authentication,
		pipeline.TenantScope,
		pipeline.AccessContext,
		pipeline.Authorization,
		pipeline.ModuleDeployment,
	} {
		if item != nil {
			items = append(items, item)
		}
	}
	return items
}
