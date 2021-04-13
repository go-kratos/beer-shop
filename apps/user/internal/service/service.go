package service

import (
	"github.com/go-kratos/beer-shop/apps/user/api/user/v1"
	"github.com/go-kratos/beer-shop/apps/user/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	v1.UnimplementedUserServer

	uc  *biz.UserUseCase
	ac  *biz.AddressUseCase
	cc  *biz.CardUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, cc *biz.CardUseCase, ac *biz.AddressUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		ac:  ac,
		cc:  cc,
		log: log.NewHelper("service/user", logger)}
}
