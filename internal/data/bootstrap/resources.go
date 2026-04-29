package bootstrap

import "github.com/chnxq/xkitpkg/app"

type DataResources struct{}

func NewDataResources(appCtx *app.AppCtx) (*DataResources, func(), error) {
	_ = appCtx
	return &DataResources{}, func() {}, nil
}
