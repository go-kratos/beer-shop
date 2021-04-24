package data

import (
	"context"
	"errors"
	"github.com/go-kratos/beer-shop/app/shop/interface/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	usV1 "github.com/go-kratos/beer-shop/api/user/service/v1"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper("repo/user", logger),
	}
}

func (rp *userRepo) Register(ctx context.Context, u *biz.User) (*biz.User, error) {
	reply, err := rp.data.uc.CreateUser(ctx, &usV1.CreateUserReq{
		Username: u.Username,
		Password: u.Password,
	})

	return &biz.User{
		Id:       reply.Id,
		Username: reply.Username,
	}, err
}

func (rp *userRepo) Login(ctx context.Context, u *biz.User) (string, error) {
	reply, err := rp.data.uc.VerifyPassword(ctx, &usV1.VerifyPasswordReq{
		Username: u.Username,
		Password: u.Password,
	})
	if err != nil {
		return "", err
	}
	if reply.Ok {
		return "some_token", nil
	}
	return "", errors.New("login failed")
}

func (rp *userRepo) Logout(ctx context.Context, u *biz.User) error {
	return nil
}
