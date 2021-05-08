// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/beer-shop/app/cart/service/internal/biz"
	"github.com/go-kratos/beer-shop/app/cart/service/internal/conf"
	"github.com/go-kratos/beer-shop/app/cart/service/internal/data"
	"github.com/go-kratos/beer-shop/app/cart/service/internal/server"
	"github.com/go-kratos/beer-shop/app/cart/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
