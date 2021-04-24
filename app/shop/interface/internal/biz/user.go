package biz

import (
	"context"
	"errors"
	csV1 "github.com/go-kratos/beer-shop/api/cart/service/v1"
	usV1 "github.com/go-kratos/beer-shop/api/user/service/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id       int64
	Username string
	Password string
}

type UserRepo interface {
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper

	us usV1.UserClient
	cs csV1.CartClient
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

func (uc *UserUseCase) Register(ctx context.Context, u *User) (*User, error) {
	reply, err := uc.us.CreateUser(ctx, &usV1.CreateUserReq{
		Username: u.Username,
		Password: u.Password,
	})

	return &User{
		Id:       reply.Id,
		Username: reply.Username,
	}, err
}

func (uc *UserUseCase) Login(ctx context.Context, u *User) (string, error) {
	reply, err := uc.us.VerifyPassword(ctx, &usV1.VerifyPasswordReq{
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

func (uc *UserUseCase) Logout(ctx context.Context, u *User) error {
	return nil
}
