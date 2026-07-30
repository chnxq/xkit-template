package server

import (
	"context"
	"reflect"
	"testing"

	"github.com/chnxq/xkitpkg/middleware"
)

func TestAppendAccessPipelinePreservesSecurityOrder(t *testing.T) {
	var order []string
	stage := func(name string) middleware.Middleware {
		return func(next middleware.Handler) middleware.Handler {
			return func(ctx context.Context, req any) (any, error) {
				order = append(order, name)
				return next(ctx, req)
			}
		}
	}
	pipeline := AccessPipeline{
		Authentication:   stage("authentication"),
		TenantScope:      stage("tenant-scope"),
		AccessContext:    stage("access-context"),
		Authorization:    stage("authorization"),
		ModuleDeployment: stage("module-deployment"),
	}
	items := appendAccessPipeline(nil, pipeline)
	handler := middleware.Chain(items...)(func(context.Context, any) (any, error) {
		order = append(order, "handler")
		return nil, nil
	})
	if _, err := handler(context.Background(), nil); err != nil {
		t.Fatal(err)
	}
	want := []string{"authentication", "tenant-scope", "access-context", "authorization", "module-deployment", "handler"}
	if !reflect.DeepEqual(order, want) {
		t.Fatalf("pipeline order = %v, want %v", order, want)
	}
}
