package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/beer-shop/app/cart/service/internal/biz"
)

var _ biz.CartRepo = (*cartRepo)(nil)

type cartRepo struct {
	data *Data
	log  *log.Helper
}

func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {
	return &cartRepo{
		data: data,
		log:  log.NewHelper("data/beer", logger),
	}
}

func (r *cartRepo)	GetCart(ctx context.Context, uid int64) (*biz.Cart, error) {
	return nil, nil
}
func (r *cartRepo)	DeleteCart(ctx context.Context, uid int64) (*biz.Cart, error) {
	return nil, nil
}
func (r *cartRepo)	AddItem(ctx context.Context, uid int64, item *biz.Item) (*biz.Cart, error) {
	return nil, nil
}
func (r *cartRepo)	UpdateItem(ctx context.Context, uid int64, item *biz.Item) (*biz.Cart, error) {
	return nil, nil
}
func (r *cartRepo)	DeleteItem(ctx context.Context, uid int64, item *biz.Item) (*biz.Cart, error) {
	return nil, nil
}

