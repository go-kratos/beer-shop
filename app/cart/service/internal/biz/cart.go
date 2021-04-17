package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Item struct {
	Id       int64
	Quantity int64
}

type Cart struct {
	Uid   int64
	Items []Item
}

type CartRepo interface {
	GetCart(ctx context.Context, uid int64) (*Cart, error)
	DeleteCart(ctx context.Context, uid int64) (*Cart, error)

	AddItem(ctx context.Context, uid int64, item *Item) (*Cart, error)
	UpdateItem(ctx context.Context, uid int64, item *Item) (*Cart, error)
	DeleteItem(ctx context.Context, uid int64, item *Item) (*Cart, error)
}

type CartUseCase struct {
	repo CartRepo
	log  *log.Helper
}

func NewCartUseCase(repo CartRepo, logger log.Logger) *CartUseCase {
	return &CartUseCase{repo: repo, log: log.NewHelper("usecase/beer", logger)}
}

func (uc *CartUseCase) GetCart(ctx context.Context, uid int64) (*Cart, error) {
	return uc.repo.GetCart(ctx, uid)
}

func (uc *CartUseCase) DeleteCart(ctx context.Context, uid int64) (*Cart, error) {
	return uc.repo.DeleteCart(ctx, uid)
}

func (uc *CartUseCase) AddItem(ctx context.Context, uid int64) (*Cart, error) {
	return uc.repo.DeleteCart(ctx, uid)
}

func (uc *CartUseCase) UpdateItem(ctx context.Context, uid int64) (*Cart, error) {
	return uc.repo.DeleteCart(ctx, uid)
}

func (uc *CartUseCase) DeleteItem(ctx context.Context, uid int64) (*Cart, error) {
	return uc.repo.DeleteCart(ctx, uid)
}
