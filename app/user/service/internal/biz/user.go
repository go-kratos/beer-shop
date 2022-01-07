package biz

import (
	"context"
	"errors"
	"math/rand"

	v1 "github.com/go-kratos/beer-shop/api/user/service/v1"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	Id       int64
	Username string
	Password string
}

type UserRepo interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	GetUser(ctx context.Context, id int64) (*User, error)
	VerifyPassword(ctx context.Context, u *User) (bool, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/user"))}
}

func (uc *UserUseCase) Save(ctx context.Context, in *v1.SaveUserReq) (*v1.SaveUserReply, error) {
	user := &User{
		Id:       rand.Int63(),
		Username: in.Username,
		Password: in.Password,
	}
	_, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		// todo: handle error
		return nil, err
	}
	return &v1.SaveUserReply{
		Id: user.Id,
	}, nil
}

func (uc *UserUseCase) GetUserByUsername(ctx context.Context, in *v1.GetUserByUsernameReq) (*v1.GetUserByUsernameReply, error) {
	user, err := uc.repo.FindByUsername(ctx, in.Username)
	if err != nil {
		//todo: handle error
		return nil, err
	}
	return &v1.GetUserByUsernameReply{
		Id:       user.Id,
		Username: user.Username,
	}, nil
}

func (uc *UserUseCase) Create(ctx context.Context, u *User) (*User, error) {
	out, err := uc.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (uc *UserUseCase) Get(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUseCase) VerifyPassword(ctx context.Context, u *User) (bool, error) {
	return uc.repo.VerifyPassword(ctx, u)
}
