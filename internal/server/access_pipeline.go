package server

import "github.com/chnxq/xkitpkg/middleware"

type AccessPipeline struct {
	Authentication   middleware.Middleware
	TenantScope      middleware.Middleware
	AccessContext    middleware.Middleware
	Authorization    middleware.Middleware
	ModuleDeployment middleware.Middleware
}

// accessPipelineBuilder is supplied by the host when it has authentication
// and authorization implementations. Generic generated projects may leave it nil.
var accessPipelineBuilder func(GeneratedData) AccessPipeline

func BuildAccessPipeline(data GeneratedData) AccessPipeline {
	if accessPipelineBuilder != nil {
		return accessPipelineBuilder(data)
	}
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
