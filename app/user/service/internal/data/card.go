package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/beer-shop/app/user/service/internal/biz"
)

var _ biz.CardRepo = (*cardRepo)(nil)

type cardRepo struct {
	data *Data
	log  *log.Helper
}

func NewCardRepo(data *Data, logger log.Logger) biz.CardRepo {
	return &cardRepo{
		data: data,
		log:  log.NewHelper("data/card", logger),
	}
}

func (r *cardRepo) CreateCard(ctx context.Context, a *biz.Card) (*biz.Card, error) {
	po, err := r.data.db.Card.
		Create().
		SetCardNo(a.CardNo).
		SetCcv(a.CCV).
		SetExpires(a.Expires).
		Save(ctx)
	return &biz.Card{
		Id:      po.ID,
		CardNo:  po.CardNo,
		CCV:     po.Ccv,
		Expires: po.Expires,
	}, err
}

func (r *cardRepo) GetCard(ctx context.Context, id int64) (*biz.Card, error) {
	po, err := r.data.db.Card.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Card{
		Id:      po.ID,
		CardNo:  po.CardNo,
		CCV:     po.Ccv,
		Expires: po.Expires,
	}, err
}

func (r *cardRepo) ListCard(ctx context.Context, uid int64) ([]*biz.Card, error) {
	pos, err := r.data.db.Card.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Card, 0)
	for _, po := range pos {
		rv = append(rv, &biz.Card{
			Id:      po.ID,
			CardNo:  po.CardNo,
			CCV:     po.Ccv,
			Expires: po.Expires,
		})
	}
	return rv, err
}
