package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"math/rand"
)

var (
	ErrPasswordInvalid = errors.New("password invalid")
	ErrUsernameInvalid = errors.New("username invalid")
	ErrUserNotFound    = errors.New("user not found")
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

func NewUser(
	username string,
	password string,
) (User, error) {

	// check username
	if len(username) <= 0 {
		return User{}, ErrUsernameInvalid
	}
	// check password
	if len(password) <= 0 {
		return User{}, ErrPasswordInvalid
	}

	return User{
		Id:       rand.Int63(),
		Username: username,
		Password: password,
	}, nil
}

type UserRepo interface {
	Find(ctx context.Context, id int64) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	Save(ctx context.Context, u *User) error

	VerifyPassword(ctx context.Context, u *User, password string) error

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
