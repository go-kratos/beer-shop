package biz

import (
	csV1 "github.com/go-kratos/beer-shop/api/cart/service/v1"
	usV1 "github.com/go-kratos/beer-shop/api/user/service/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
}

type UserRepo interface {
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper

	us   usV1.UserClient
	cs   csV1.CartClient
}

func NewUserUseCase(repo UserRepo, logger log.Logger, us usV1.UserClient, cs csV1.CartClient) *UserUseCase {
	log := log.NewHelper("usecase/interface", logger)
	return &UserUseCase{
		repo: repo,
		us:   us,
		cs:   cs,
		log:  log,
	}
}
