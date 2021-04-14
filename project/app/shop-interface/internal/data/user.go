package data

import (
	"github.com/go-kratos/beer-shop/project/app/shop-interface/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper("data/server-service", logger),
	}
}
