// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/beer-shop/app/shop/interface/internal/biz"
	"github.com/go-kratos/beer-shop/app/shop/interface/internal/conf"
	"github.com/go-kratos/beer-shop/app/shop/interface/internal/data"
	"github.com/go-kratos/beer-shop/app/shop/interface/internal/server"
	"github.com/go-kratos/beer-shop/app/shop/interface/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, auth *conf.Auth, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	discovery := data.NewDiscovery(registry)
	userClient := data.NewUserServiceClient(discovery, tracerProvider)
	cartClient := data.NewCartServiceClient(discovery, tracerProvider)
	catalogClient := data.NewCatalogServiceClient(discovery, tracerProvider)
	dataData, err := data.NewData(confData, logger, userClient, cartClient, catalogClient)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, logger)
	beerRepo := data.NewBeerRepo(dataData, logger)
	catalogUseCase := biz.NewCatalogUseCase(beerRepo, logger)
	shopInterface := service.NewShopInterface(userUseCase, catalogUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, logger, tracerProvider, shopInterface)
	grpcServer := server.NewGRPCServer(confServer, logger, tracerProvider, shopInterface)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
	}, nil
}
