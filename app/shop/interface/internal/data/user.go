package data

import (
	"context"
	"errors"

	"github.com/go-kratos/beer-shop/app/shop/interface/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	usV1 "github.com/go-kratos/beer-shop/app/shop/interface/internal/api/user/service/v1"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
	}
}

func (rp *userRepo) Register(ctx context.Context, u *biz.User) (*biz.User, error) {
	reply, err := rp.data.uc.CreateUser(ctx, &usV1.CreateUserReq{
		Username: u.Username,
		Password: u.Password,
	})
	if err != nil {
		return nil, err
	}

	return &biz.User{
		Id:       reply.Id,
		Username: reply.Username,
	}, err
}

func (rp *userRepo) Login(ctx context.Context, u *biz.User) (string, error) {
	reply, err := rp.data.uc.VerifyPassword(ctx, &usV1.VerifyPasswordReq{
		Username: u.Username,
		Password: u.Password,
	})
	if err != nil {
		return "", err
	}
	if reply.Ok {
		return "some_token", nil
	}
	return "", errors.New("login failed")
}

func (rp *userRepo) Logout(ctx context.Context, u *biz.User) error {
	return nil
}

func (rp *userRepo) CreateAddress(ctx context.Context, uid int64, a *biz.Address) (*biz.Address, error) {
	reply, err := rp.data.uc.CreateAddress(ctx, &usV1.CreateAddressReq{
		Uid:      uid,
		Name:     a.Name,
		Mobile:   a.Mobile,
		Address:  a.Address,
		PostCode: a.PostCode,
	})
	if err != nil {
		return nil, err
	}

	return &biz.Address{
		Id:       reply.Id,
		Name:     reply.Name,
		Mobile:   reply.Mobile,
		Address:  reply.Address,
		PostCode: reply.PostCode,
	}, err
}

func (rp *userRepo) GetAddress(ctx context.Context, id int64) (*biz.Address, error) {
	reply, err := rp.data.uc.GetAddress(ctx, &usV1.GetAddressReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return &biz.Address{
		Id:       reply.Id,
		Name:     reply.Name,
		Mobile:   reply.Mobile,
		Address:  reply.Address,
		PostCode: reply.PostCode,
	}, err
}

func (rp *userRepo) ListAddress(ctx context.Context, uid int64) ([]*biz.Address, error) {
	reply, err := rp.data.uc.ListAddress(ctx, &usV1.ListAddressReq{
		Uid: uid,
	})
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Address, 0)
	for _, x := range reply.Results {
		rv = append(rv, &biz.Address{
			Id:       x.Id,
			Name:     x.Name,
			Mobile:   x.Mobile,
			Address:  x.Address,
			PostCode: x.PostCode,
		})
	}
	return rv, err
}

func (rp *userRepo) CreateCard(ctx context.Context, uid int64, c *biz.Card) (*biz.Card, error) {
	reply, err := rp.data.uc.CreateCard(ctx, &usV1.CreateCardReq{
		Uid:     uid,
		CardNo:  c.CardNo,
		Ccv:     c.CCV,
		Expires: c.Expires,
	})
	if err != nil {
		return nil, err
	}

	return &biz.Card{
		Id: reply.Id,
	}, err
}

func (rp *userRepo) GetCard(ctx context.Context, id int64) (*biz.Card, error) {
	reply, err := rp.data.uc.GetCard(ctx, &usV1.GetCardReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return &biz.Card{
		Id:     reply.Id,
		CardNo: reply.CardNo,
	}, err
}

func (rp *userRepo) ListCard(ctx context.Context, uid int64) ([]*biz.Card, error) {
	reply, err := rp.data.uc.ListCard(ctx, &usV1.ListCardReq{
		Uid: uid,
	})
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Card, 0)
	for _, x := range reply.Results {
		rv = append(rv, &biz.Card{
			Id:     x.Id,
			CardNo: x.CardNo,
		})
	}
	return rv, err
}
