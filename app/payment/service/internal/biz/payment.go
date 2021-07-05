package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type PaymentRepo interface {
}

type PaymentUseCase struct {
	repo PaymentRepo
	log  *log.Helper
}

func NewPaymentUseCase(repo PaymentRepo, logger log.Logger) *PaymentUseCase {
	return &PaymentUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/payment"))}
}
