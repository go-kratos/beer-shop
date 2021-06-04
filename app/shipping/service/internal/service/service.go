package service

import (
	v1 "github.com/go-kratos/beer-shop/api/shipping/service/v1"
	"github.com/go-kratos/beer-shop/app/shipping/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewShippingService)

type ShippingService struct {
	v1.UnimplementedShippingServer

	oc  *biz.ShippingUseCase
	log *log.Helper
}

func NewShippingService(oc *biz.ShippingUseCase, logger log.Logger) *ShippingService {
	return &ShippingService{

		oc:  oc,
		log: log.NewHelper(log.With(logger, "module", "service/shipping"))}
}
