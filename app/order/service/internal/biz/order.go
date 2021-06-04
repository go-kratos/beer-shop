package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Order struct {
	Id     int64
	UserId int64
}

type OrderRepo interface {
	CreateOrder(ctx context.Context, c *Order) (*Order, error)
	GetOrder(ctx context.Context, id int64) (*Order, error)
	UpdateOrder(ctx context.Context, c *Order) (*Order, error)
	ListOrder(ctx context.Context, pageNum, pageSize int64) ([]*Order, error)
}

type OrderUseCase struct {
	repo OrderRepo
	log  *log.Helper
}

func NewOrderUseCase(repo OrderRepo, logger log.Logger) *OrderUseCase {
	return &OrderUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/order"))}
}

func (uc *OrderUseCase) Create(ctx context.Context, u *Order) (*Order, error) {
	return uc.repo.CreateOrder(ctx, u)
}

func (uc *OrderUseCase) Get(ctx context.Context, id int64) (*Order, error) {
	return uc.repo.GetOrder(ctx, id)
}

func (uc *OrderUseCase) Update(ctx context.Context, u *Order) (*Order, error) {
	return uc.repo.UpdateOrder(ctx, u)
}

func (uc *OrderUseCase) List(ctx context.Context, pageNum, pageSize int64) ([]*Order, error) {
	return uc.repo.ListOrder(ctx, pageNum, pageSize)
}
