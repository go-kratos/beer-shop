package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id       int64
	Username string
	Password string
}

type UserRepo interface {
	Register(ctx context.Context, u *User) (*User, error)
	Login(ctx context.Context, u *User) (string, error)
}

type UserUseCase struct {
	repo   UserRepo
	log    *log.Helper
	authUc *AuthUseCase
}

func NewUserUseCase(repo UserRepo, logger log.Logger, authUc *AuthUseCase) *UserUseCase {
	log := log.NewHelper(log.With(logger, "module", "usecase/interface"))
	return &UserUseCase{
		repo:   repo,
		log:    log,
		authUc: authUc,
	}
}

func (uc *UserUseCase) Register(ctx context.Context, u *User) (*User, error) {
	return uc.repo.Register(ctx, u)
}

func (uc *UserUseCase) Login(ctx context.Context, u *User) (string, error) {
	token, err := uc.repo.Login(ctx, u)
	if err != nil {
		return token, err
	}
	return uc.authUc.Auth(u.Id)
}

func (uc *UserUseCase) Logout(ctx context.Context, u *User) error {
	return nil
}
