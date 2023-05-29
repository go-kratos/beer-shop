package data

import (
	"context"

	"github.com/go-kratos/beer-shop/app/user/service/internal/biz"
	"github.com/go-kratos/beer-shop/app/user/service/internal/data/ent/user"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.CardRepo = (*cardRepo)(nil)

type cardRepo struct {
	data *Data
	log  *log.Helper
}

func NewCardRepo(data *Data, logger log.Logger) biz.CardRepo {
	return &cardRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/card")),
	}
}

func (r *cardRepo) CreateCard(ctx context.Context, c *biz.Card) (*biz.Card, error) {
	po, err := r.data.db.Card.
		Create().
		SetCardNo(c.CardNo).
		SetCcv(c.CCV).
		SetExpires(c.Expires).
		SetName(c.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Card{
		Id:      po.ID,
		CardNo:  po.CardNo,
		CCV:     po.Ccv,
		Expires: po.Expires,
		Name:    po.Name,
	}, nil
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
	}, nil
}

func (r *cardRepo) ListCard(ctx context.Context, uid int64) ([]*biz.Card, error) {
	pos, err := r.data.db.User.
		Query().
		Where(user.ID(uid)).
		QueryCards().
		All(ctx)
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
	return rv, nil
}
