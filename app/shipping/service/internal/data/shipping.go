package data

import (
	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/beer-shop/app/shipping/service/internal/biz"
)

var _ biz.ShippingRepo = (*shippingRepo)(nil)

type shippingRepo struct {
	data *Data
	log  *log.Helper
}

func NewShippingRepo(data *Data, logger log.Logger) biz.ShippingRepo {
	return &shippingRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/shipping")),
	}
}
