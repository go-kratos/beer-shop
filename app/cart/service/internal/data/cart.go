package data

import (
	"context"
	"github.com/go-kratos/beer-shop/pkg/utils/pagination"

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

func (r *cartRepo) CreateBeer(ctx context.Context, b *biz.Beer) (*biz.Beer, error) {
	po, err := r.data.db.Beer.
		Create().
		SetName(b.Name).
		SetDescription(b.Description).
		SetCount(b.Count).
		SetImages(b.Images).
		Save(ctx)
	return &biz.Beer{
		Id:          po.ID,
		Description: po.Description,
		Count:       po.Count,
		Images:      po.Images,
	}, err
}

func (r *cartRepo) GetBeer(ctx context.Context, id int64) (*biz.Beer, error) {
	po, err := r.data.db.Beer.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Beer{
		Id:          po.ID,
		Description: po.Description,
		Count:       po.Count,
		Images:      po.Images,
	}, err
}

func (r *cartRepo) ListBeer(ctx context.Context, pageNum, pageSize int64) ([]*biz.Beer, error) {
	pos, err := r.data.db.Beer.Query().
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Limit(int(pageSize)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Beer, 0)
	for _, po := range pos {
		rv = append(rv, &biz.Beer{
			Id:          po.ID,
			Description: po.Description,
			Count:       po.Count,
			Images:      po.Images,
		})
	}
	return rv, err
}
