package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Username string
	Password string
}

type UserRepo interface {
	CreateUser(*User) (*User, error)
	GetUser(*User) (*User, error)
	VerifyPassword(*User) (bool, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper("usecase/user", logger)}
}

func (uc *UserUsecase) Create(u *User) (*User, error) {
	return uc.repo.CreateUser(u)
}

func (uc *UserUsecase) Get(u *User) (*User, error) {
	return uc.repo.GetUser(u)
}

func (uc *UserUsecase) VerifyPassword(u *User) (bool, error) {
	return uc.repo.VerifyPassword(u)
}
