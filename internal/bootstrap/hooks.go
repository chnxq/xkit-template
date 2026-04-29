package bootstrap

import (
	"github.com/chnxq/xkitpkg/app"
	"github.com/chnxq/xkitpkg/transport"

	"xkit-template-v01/internal/server"
)

func NewManualServers(appCtx *app.AppCtx) ([]transport.Server, func(), error) {
	servers := make([]transport.Server, 0, 2)

	asynqServer, err := server.NewAsynqServer(appCtx)
	if err != nil {
		return nil, func() {}, err
	}
	if asynqServer != nil {
		servers = append(servers, asynqServer)
	}

	sseServer, err := server.NewSSEServer(appCtx)
	if err != nil {
		return nil, func() {}, err
	}
	if sseServer != nil {
		servers = append(servers, sseServer)
	}

	return servers, func() {}, nil
}
