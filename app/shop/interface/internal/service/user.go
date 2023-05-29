package service

import (
	"context"

	v1 "github.com/go-kratos/beer-shop/api/shop/interface/v1"
	"github.com/go-kratos/beer-shop/app/shop/interface/internal/biz"
)

func (s *ShopInterface) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	return s.ac.Register(ctx, req)
}

func (s *ShopInterface) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	return s.ac.Login(ctx, req)
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
	if err != nil {
		return nil, err
	}
	return &v1.CreateAddressReply{
		Id: rv.Id,
	}, nil
}

func (s *ShopInterface) GetAddress(ctx context.Context, req *v1.GetAddressReq) (*v1.GetAddressReply, error) {
	x, err := s.uc.GetAddress(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.GetAddressReply{
		Id:       x.Id,
		Name:     x.Name,
		Mobile:   x.Mobile,
		Address:  x.Address,
		PostCode: x.PostCode,
	}, nil
}

func (s *ShopInterface) ListCard(ctx context.Context, req *v1.ListCardReq) (*v1.ListCardReply, error) {
	rv, err := s.uc.ListCard(ctx, req.Uid)
	if err != nil {
		return nil, err
	}

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
	}, nil
}

func (s *ShopInterface) CreateCard(ctx context.Context, req *v1.CreateCardReq) (*v1.CreateCardReply, error) {
	rv, err := s.uc.CreateCard(ctx, req.Uid, &biz.Card{
		CardNo:  req.CardNo,
		CCV:     req.Ccv,
		Expires: req.Expires,
		Name:    req.Name,
	})
	if err != nil {
		return nil, err
	}

	return &v1.CreateCardReply{
		Id: rv.Id,
	}, nil
}

func (s *ShopInterface) GetCard(ctx context.Context, req *v1.GetCardReq) (*v1.GetCardReply, error) {
	rv, err := s.uc.GetCard(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.GetCardReply{
		Id:      rv.Id,
		CardNo:  rv.CardNo,
		Ccv:     rv.CCV,
		Expires: rv.Expires,
	}, nil
}

func (s *ShopInterface) DeleteCard(ctx context.Context, req *v1.DeleteCardReq) (*v1.DeleteCardReply, error) {
	return nil, nil
}
