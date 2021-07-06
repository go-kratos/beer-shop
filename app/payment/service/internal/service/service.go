package service

import (
	v1 "github.com/go-kratos/beer-shop/api/payment/service/v1"
	"github.com/go-kratos/beer-shop/app/payment/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewPaymentService)

type PaymentService struct {
	v1.UnimplementedPaymentServer

	pc  *biz.PaymentUseCase
	log *log.Helper
}

func NewPaymentService(pc *biz.PaymentUseCase, logger log.Logger) *PaymentService {
	return &PaymentService{
		pc:  pc,
		log: log.NewHelper(log.With(logger, "module", "service/payment"))}
}
