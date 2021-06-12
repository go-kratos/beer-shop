package server

import (
	"github.com/go-kratos/beer-shop/app/order/service/internal/conf"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"

	consul "github.com/go-kratos/consul/registry"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewRegistrar)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli)
	return r
}
