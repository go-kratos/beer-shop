package service

import (
	"github.com/go-kratos/beer-shop/api/payment/service/v1"
	"github.com/go-kratos/beer-shop/app/payment/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewPaymentService)

type PaymentService struct {
	v1.UnimplementedPaymentServer

	bc  *biz.BeerUseCase
	log *log.Helper
}

func NewPaymentService(bc *biz.BeerUseCase, logger log.Logger) *PaymentService {
	return &PaymentService{

		bc:  bc,
		log: log.NewHelper(log.With(logger, "module", "service/payment"))}
}
