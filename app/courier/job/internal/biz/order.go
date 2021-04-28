package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type Courier struct {
	Id     int64
	UserId int64
}

type CourierRepo interface {
}

type CourierUseCase struct {
	repo CourierRepo
	log  *log.Helper
}

func NewCourierUseCase(repo CourierRepo, logger log.Logger) *CourierUseCase {
	return &CourierUseCase{repo: repo, log: log.NewHelper("usecase/courier", logger)}
}
