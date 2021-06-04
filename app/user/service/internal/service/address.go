package service

import (
	"context"

	v1 "github.com/go-kratos/beer-shop/api/user/service/v1"
	"github.com/go-kratos/beer-shop/app/user/service/internal/biz"
)

func (s *UserService) CreateAddress(ctx context.Context, req *v1.CreateAddressReq) (*v1.CreateAddressReply, error) {
	rv, err := s.ac.Create(ctx, req.Uid, &biz.Address{
		Name:     req.Name,
		Mobile:   req.Mobile,
		Address:  req.Address,
		PostCode: req.PostCode,
	})
	return &v1.CreateAddressReply{
		Id: rv.Id,
	}, err
}

func (s *UserService) GetAddress(ctx context.Context, req *v1.GetAddressReq) (*v1.GetAddressReply, error) {
	x, err := s.ac.Get(ctx, req.Id)
	return &v1.GetAddressReply{
		Id:       x.Id,
		Name:     x.Name,
		Mobile:   x.Mobile,
		Address:  x.Address,
		PostCode: x.PostCode,
	}, err
}

func (s *UserService) ListAddress(ctx context.Context, req *v1.ListAddressReq) (*v1.ListAddressReply, error) {
	rv, err := s.ac.List(ctx, req.Uid)
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
