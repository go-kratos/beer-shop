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
	UserId int64
	Items  []Item
}

type CartRepo interface {
	GetCart(ctx context.Context, uid int64) (*Cart, error)
	SaveCart(ctx context.Context, c *Cart) error
	DeleteCart(ctx context.Context, uid int64) error
}

type CartUseCase struct {
	repo CartRepo
	log  *log.Helper
}

func NewCartUseCase(repo CartRepo, logger log.Logger) *CartUseCase {
	return &CartUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/cart"))}
}

func (uc *CartUseCase) GetCart(ctx context.Context, uid int64) (*Cart, error) {
	return uc.repo.GetCart(ctx, uid)
}

func (uc *CartUseCase) DeleteCart(ctx context.Context, uid int64) error {
	return uc.repo.DeleteCart(ctx, uid)
}

func (uc *CartUseCase) AddItem(ctx context.Context, uid int64, in Item) (c *Cart, err error) {
	if c, err = uc.repo.GetCart(ctx, uid); err != nil {
		return nil, err
	}
	hit := false
	for _, x := range c.Items {
		if x.Id == in.Id {
			x.Quantity += in.Quantity
			hit = true
			break
		}
	}
	if !hit {
		c.Items = append(c.Items, in)
	}
	if err := uc.repo.SaveCart(ctx, c); err != nil {
		return nil, err
	}
	return uc.repo.GetCart(ctx, uid)
}

func (uc *CartUseCase) UpdateItem(ctx context.Context, uid, itemId, quantity int64) (c *Cart, err error) {
	if c, err = uc.repo.GetCart(ctx, uid); err != nil {
		return nil, err
	}
	hit := false
	for _, x := range c.Items {
		if x.Id == itemId {
			x.Quantity = quantity
			hit = true
			break
		}
	}
	if !hit {
		c.Items = append(c.Items, Item{Id: itemId, Quantity: quantity})
	}
	if err := uc.repo.SaveCart(ctx, c); err != nil {
		return nil, err
	}
	return uc.repo.GetCart(ctx, uid)
}

func (uc *CartUseCase) DeleteItem(ctx context.Context, uid int64, itemId int64) (c *Cart, err error) {
	if c, err = uc.repo.GetCart(ctx, uid); err != nil {
		return nil, err
	}
	for i, x := range c.Items {
		if x.Id == itemId {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			break
		}
	}
	if err := uc.repo.SaveCart(ctx, c); err != nil {
		return nil, err
	}
	return uc.repo.GetCart(ctx, uid)
}
