package biz

import (
	usV1 "github.com/go-kratos/beer-shop/api/user/service/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
}

type UserRepo interface {
}

type UserUseCase struct {
	repo UserRepo
	us   usV1.UserClient
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger, us usV1.UserClient) *UserUseCase {
	log := log.NewHelper(log.With(logger, "module", "usecase/interface"))
	return &UserUseCase{
		repo: repo,
		us:   us,
		log:  log,
	}
}
