package service

import (
	"context"
	"github.com/go-kratos/beer-shop/app/shop/interface/internal/api/shop/interface/v1"
	"github.com/go-kratos/beer-shop/app/shop/interface/internal/biz"
)

func (s *ShopInterface) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	rv, err := s.uc.Register(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{
		Id: rv.Id,
	}, err
}

func (s *ShopInterface) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	rv, err := s.uc.Login(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	return &v1.LoginReply{
		Token: rv,
	}, err
}

func (s *ShopInterface) Logout(ctx context.Context, req *v1.LogoutReq) (*v1.LogoutReply, error) {
	err := s.uc.Logout(ctx, &biz.User{})
	return &v1.LogoutReply{}, err
}

func (s *ShopInterface) ListAddress(ctx context.Context, req *v1.ListAddressReq) (*v1.ListAddressReply, error) {
	rv, err := s.uc.ListAddress(ctx, req.Uid)
	rs := make([]*v1.ListAddressReply_Address, 0)
	for _, x := range rv {
		rs = append(rs, &v1.ListAddressReply_Address{
			Id:       x.Id,
			Name:     x.Name,
			Mobile:   x.Mobile,
			Address:  x.Address,
			PostCode: x.PostCode,
		})
	}
	return &v1.ListAddressReply{
		Results: rs,
	}, err
}

func (s *ShopInterface) CreateAddress(ctx context.Context, req *v1.CreateAddressReq) (*v1.CreateAddressReply, error) {
	rv, err := s.uc.CreateAddress(ctx, req.Uid, &biz.Address{
		Name:     req.Name,
		Mobile:   req.Mobile,
		Address:  req.Address,
		PostCode: req.PostCode,
	})
	return &v1.CreateAddressReply{
		Id: rv.Id,
	}, err
}

func (s *ShopInterface) GetAddress(ctx context.Context, req *v1.GetAddressReq) (*v1.GetAddressReply, error) {
	x, err := s.uc.GetAddress(ctx, req.Id)
	return &v1.GetAddressReply{
		Id:       x.Id,
		Name:     x.Name,
		Mobile:   x.Mobile,
		Address:  x.Address,
		PostCode: x.PostCode,
	}, err
}

func (s *ShopInterface) ListCard(ctx context.Context, req *v1.ListCardReq) (*v1.ListCardReply, error) {
	rv, err := s.uc.ListCard(ctx, req.Uid)
	rs := make([]*v1.ListCardReply_Card, 0)
	for _, x := range rv {
		rs = append(rs, &v1.ListCardReply_Card{
			Id:      x.Id,
			CardNo:  x.CardNo,
			Ccv:     x.CCV,
			Expires: x.Expires,
		})
	}
	return &v1.ListCardReply{
		Results: rs,
	}, err
}

func (s *ShopInterface) CreateCard(ctx context.Context, req *v1.CreateCardReq) (*v1.CreateCardReply, error) {
	rv, err := s.uc.CreateCard(ctx, req.Uid, &biz.Card{
		CardNo:  req.CardNo,
		CCV:     req.Ccv,
		Expires: req.Expires,
	})
	return &v1.CreateCardReply{
		Id: rv.Id,
	}, err
}

func (s *ShopInterface) GetCard(ctx context.Context, req *v1.GetCardReq) (*v1.GetCardReply, error) {
	rv, err := s.uc.GetCard(ctx, req.Id)
	return &v1.GetCardReply{
		Id:      rv.Id,
		CardNo:  rv.CardNo,
		Ccv:     rv.CCV,
		Expires: rv.Expires,
	}, err
}

func (s *ShopInterface) DeleteCard(ctx context.Context, req *v1.DeleteCardReq) (*v1.DeleteCardReply, error) {
	return nil, nil
}
