package data

import (
	"context"
	"time"

	"github.com/go-kratos/beer-shop/pkg/util/pagination"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/beer-shop/app/order/service/internal/biz"
)

var _ biz.OrderRepo = (*orderRepo)(nil)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

type Order struct {
	gorm.Model
	Id        int64
	UserId    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/order")),
	}
}

func (r *orderRepo) CreateOrder(ctx context.Context, b *biz.Order) (*biz.Order, error) {
	o := Order{Id: b.Id, UserId: b.UserId}
	result := r.data.db.WithContext(ctx).Create(o)
	return &biz.Order{
		Id: o.Id,
	}, result.Error
}

func (r *orderRepo) GetOrder(ctx context.Context, id int64) (*biz.Order, error) {
	o := Order{}
	result := r.data.db.WithContext(ctx).First(&o, id)
	return &biz.Order{
		Id: o.Id,
	}, result.Error
}

func (r *orderRepo) UpdateOrder(ctx context.Context, b *biz.Order) (*biz.Order, error) {
	o := Order{}
	result := r.data.db.WithContext(ctx).First(&o, b.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	o.UserId = b.UserId
	result = r.data.db.WithContext(ctx).Save(&o)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.Order{
		Id: o.Id,
	}, nil
}

func (r *orderRepo) ListOrder(ctx context.Context, pageNum, pageSize int64) ([]*biz.Order, error) {
	var os []Order
	result := r.data.db.WithContext(ctx).
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Find(&os)
	if result.Error != nil {
		return nil, result.Error
	}
	rv := make([]*biz.Order, 0)
	for _, o := range os {
		rv = append(rv, &biz.Order{
			Id: o.Id,
		})
	}
	return rv, nil
}
