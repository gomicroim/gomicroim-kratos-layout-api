package service

import (
	"context"

	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService, NewUserClient)

func NewUserClient(config *conf.Discover, discovery registry.Discovery) UserClient {
	rpcUserEndpoint := "discovery:///" + config.ServiceEndpoint.RpcUser
	conn, err := grpc.DialInsecure(context.Background(),
		grpc.WithEndpoint(rpcUserEndpoint),
		grpc.WithDiscovery(discovery),
	)
	if err != nil {
		panic(err)
	}
	// replace your grpc client init code here
	return newUserClient(conn)
}
