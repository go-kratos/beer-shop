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

type Address struct {
	Id       int64
	Name     string
	Mobile   string
	Address  string
	PostCode string
}

type Card struct {
	Id      int64
	CardNo  string
	CCV     string
	Expires string
}

type UserRepo interface {
	Register(ctx context.Context, u *User) (*User, error)
	Login(ctx context.Context, u *User) (string, error)

	CreateAddress(ctx context.Context, uid int64, a *Address) (*Address, error)
	GetAddress(ctx context.Context, id int64) (*Address, error)
	ListAddress(ctx context.Context, uid int64) ([]*Address, error)

	CreateCard(ctx context.Context, uid int64, c *Card) (*Card, error)
	GetCard(ctx context.Context, id int64) (*Card, error)
	ListCard(ctx context.Context, id int64) ([]*Card, error)
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

func (uc *UserUseCase) CreateAddress(ctx context.Context, uid int64, a *Address) (*Address, error) {
	return uc.repo.CreateAddress(ctx, uid, a)
}

func (uc *UserUseCase) GetAddress(ctx context.Context, id int64) (*Address, error) {
	return uc.repo.GetAddress(ctx, id)
}

func (uc *UserUseCase) ListAddress(ctx context.Context, uid int64) ([]*Address, error) {
	return uc.repo.ListAddress(ctx, uid)
}

func (uc *UserUseCase) CreateCard(ctx context.Context, uid int64, c *Card) (*Card, error) {
	return uc.repo.CreateCard(ctx, uid, c)
}

func (uc *UserUseCase) GetCard(ctx context.Context, id int64) (*Card, error) {
	return uc.repo.GetCard(ctx, id)
}

func (uc *UserUseCase) ListCard(ctx context.Context, uid int64) ([]*Card, error) {
	return uc.repo.ListCard(ctx, uid)
}
