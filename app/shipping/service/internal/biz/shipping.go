package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type Shipping struct {
	Id     int64
	UserId int64
}

type ShippingRepo interface {
}

type ShippingUseCase struct {
	repo ShippingRepo
	log  *log.Helper
}

func NewShippingUseCase(repo ShippingRepo, logger log.Logger) *ShippingUseCase {
	return &ShippingUseCase{repo: repo, log: log.NewHelper("usecase/shipping", logger)}
}
