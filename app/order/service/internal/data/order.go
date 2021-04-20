package data

import (
	"context"
	"github.com/go-kratos/beer-shop/pkg/utils/pagination"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/beer-shop/app/order/service/internal/biz"
)

var _ biz.OrderRepo = (*orderRepo)(nil)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper("data/order", logger),
	}
}

func (r *orderRepo) CreateOrder(ctx context.Context, b *biz.Order) (*biz.Order, error) {
	po, err := r.data.db.Order.
		Create().
		SetUserID(b.UserId).
		Save(ctx)
	return &biz.Order{
		Id: po.ID,
	}, err
}

func (r *orderRepo) GetOrder(ctx context.Context, id int64) (*biz.Order, error) {
	po, err := r.data.db.Order.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Order{
		Id: po.ID,
	}, err
}

func (r *orderRepo) UpdateOrder(ctx context.Context, b *biz.Order) (*biz.Order, error) {
	po, err := r.data.db.Order.
		Create().
		SetUserID(b.UserId).
		Save(ctx)
	return &biz.Order{
		Id: po.ID,
	}, err
}

func (r *orderRepo) ListOrder(ctx context.Context, pageNum, pageSize int64) ([]*biz.Order, error) {
	pos, err := r.data.db.Order.Query().
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Limit(int(pageSize)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Order, 0)
	for _, po := range pos {
		rv = append(rv, &biz.Order{
			Id: po.ID,
		})
	}
	return rv, err
}
