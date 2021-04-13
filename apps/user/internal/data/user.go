package data

import (
	"github.com/go-kratos/beer-shop/apps/user/internal/biz"
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
		log:  log.NewHelper("data/user", logger),
	}
}


func (r *userRepo) CreateUser(*biz.User) (*biz.User, error) {
	return nil, nil
}

func (r *userRepo) GetUser(*biz.User) (*biz.User, error) {
	return nil, nil
}

func (r *userRepo) VerifyPassword(*biz.User) (bool, error) {
	return false, nil
}
