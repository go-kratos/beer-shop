package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/beer-shop/app/user/service/internal/biz"
)

var _ biz.AddressRepo = (*addressRepo)(nil)

type addressRepo struct {
	data *Data
	log  *log.Helper
}

func NewAddressRepo(data *Data, logger log.Logger) biz.AddressRepo {
	return &addressRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/address")),
	}
}

func (r *addressRepo) CreateAddress(ctx context.Context, a *biz.Address) (*biz.Address, error) {
	po, err := r.data.db.Address.
		Create().
		SetName(a.Name).
		SetAddress(a.Address).
		SetMobile(a.Mobile).
		SetPostCode(a.PostCode).
		Save(ctx)
	return &biz.Address{
		Id:       po.ID,
		Name:     po.Name,
		Mobile:   po.Mobile,
		PostCode: po.PostCode,
		Address:  po.Address,
	}, err
}

func (r *addressRepo) GetAddress(ctx context.Context, id int64) (*biz.Address, error) {
	po, err := r.data.db.Address.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Address{
		Id:       po.ID,
		Name:     po.Name,
		Mobile:   po.Mobile,
		PostCode: po.PostCode,
		Address:  po.Address,
	}, err
}

func (r *addressRepo) ListAddress(ctx context.Context, uid int64) ([]*biz.Address, error) {
	pos, err := r.data.db.Address.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Address, 0)
	for _, po := range pos {
		rv = append(rv, &biz.Address{
			Id:       po.ID,
			Name:     po.Name,
			Mobile:   po.Mobile,
			PostCode: po.PostCode,
			Address:  po.Address,
		})
	}
	return rv, err
}
