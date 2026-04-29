package server

import (
	"fmt"

	"github.com/chnxq/xkitpkg/app"
	grpctransport "github.com/chnxq/xkitpkg/transport/grpc"
)

func NewGRPCServer(appCtx *app.AppCtx, services GeneratedGRPCServices) (*grpctransport.Server, error) {
	opts, err := GRPCServerOptions(appCtx)
	if err != nil {
		return nil, fmt.Errorf("grpc server options: %w", err)
	}
	srv := grpctransport.NewServer(opts...)
	RegisterGeneratedGRPCServices(srv, services)
	//RegisterManualGRPCServices(srv, appCtx)
	return srv, nil
}
