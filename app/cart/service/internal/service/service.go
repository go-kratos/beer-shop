package service

import (
	"github.com/go-kratos/beer-shop/api/cart/service/v1"
	"github.com/go-kratos/beer-shop/app/cart/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCartService)

type CartService struct {
	v1.UnimplementedCartServer

	bc  *biz.BeerUseCase
	log *log.Helper
}

func NewCartService(bc *biz.BeerUseCase, logger log.Logger) *CartService {
	return &CartService{

		bc:  bc,
		log: log.NewHelper("service/cart", logger)}
}
