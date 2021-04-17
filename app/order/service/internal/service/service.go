package service

import (
	"github.com/go-kratos/beer-shop/api/order/service/v1"
	"github.com/go-kratos/beer-shop/app/order/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewOrderService)

type OrderService struct {
	v1.UnimplementedOrderServer

	bc  *biz.BeerUseCase
	log *log.Helper
}

func NewOrderService(bc *biz.BeerUseCase, logger log.Logger) *OrderService {
	return &OrderService{

		bc:  bc,
		log: log.NewHelper("service/order", logger)}
}
