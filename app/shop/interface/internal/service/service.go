package service

import (
	"github.com/go-kratos/beer-shop/api/shop/interface/v1"
	"github.com/go-kratos/beer-shop/app/shop/interface/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewShopInterface)

type ShopInterface struct {
	v1.UnimplementedShopInterfaceServer

	uc *biz.UserUseCase
	ac *biz.AuthUseCase
	cc *biz.CatalogUseCase

	log *log.Helper
}

func NewShopInterface(
	uc *biz.UserUseCase,
	cc *biz.CatalogUseCase,
	ac *biz.AuthUseCase,
	logger log.Logger) *ShopInterface {
	return &ShopInterface{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
		uc:  uc,
		ac:  ac,
		cc:  cc,
	}
}
