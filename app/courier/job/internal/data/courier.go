package data

import (
	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/beer-shop/app/courier/job/internal/biz"
)

var _ biz.CourierRepo = (*courierRepo)(nil)

type courierRepo struct {
	data *Data
	log  *log.Helper
}

func NewCourierRepo(data *Data, logger log.Logger) biz.CourierRepo {
	return &courierRepo{
		data: data,
		log:  log.NewHelper("data/courier", logger),
	}
}
