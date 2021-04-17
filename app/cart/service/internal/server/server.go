package server

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"

	consul "github.com/go-kratos/consul/registry"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer, NewRegistrar)

func NewRegistrar() registry.Registrar {
	cli, err := consulAPI.NewClient(consulAPI.DefaultConfig())
	if err != nil {
		panic(err)
	}
	r := consul.New(cli)
	return r
}
