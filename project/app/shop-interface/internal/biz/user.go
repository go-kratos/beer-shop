package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {

}

type UserRepo interface {
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper("usecase/shop-interface", logger)}
}
