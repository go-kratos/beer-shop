package biz

import (
	"context"
	usV1 "github.com/go-kratos/beer-shop/project/app/user-service/api/user-service/v1"

	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUserUseCase, NewDiscovery, NewUserServiceClient)

func NewDiscovery() registry.Discovery {
	cli, err := consulAPI.NewClient(consulAPI.DefaultConfig())
	if err != nil {
		panic(err)
	}
	r := consul.New(cli)
	return r
}

func NewUserServiceClient(r registry.Discovery) usV1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///helloworld"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		panic(err)
	}
	return usV1.NewUserClient(conn)
}
