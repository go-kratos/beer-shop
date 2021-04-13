package service

import (
	"github.com/go-kratos/beer-shop/apps/user/api/user/v1"
	"github.com/go-kratos/beer-shop/apps/user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	v1.UnimplementedUserServer

	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper("service/user", logger)}
}

