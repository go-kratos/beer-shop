package data

import (
	"github.com/go-kratos/beer-shop/app/payment/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.PaymentRepo = (*paymentRepo)(nil)

type paymentRepo struct {
	data *Data
	log  *log.Helper
}

func NewPaymentRepo(data *Data, logger log.Logger) biz.PaymentRepo {
	return &paymentRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/payment")),
	}
}
