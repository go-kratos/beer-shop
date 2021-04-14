package service

import (
	"github.com/go-kratos/beer-shop/project/app/shop-interface/api/user/v1"
	"github.com/go-kratos/beer-shop/project/app/shop-interface/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewShopService)

type UserService struct {
	v1.UnimplementedUserServer

	log *log.Helper
}

func NewShopService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		log: log.NewHelper("service/server-service", logger),
	}
}
