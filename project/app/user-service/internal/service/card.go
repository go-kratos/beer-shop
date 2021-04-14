package service

import (
	"context"
	"github.com/go-kratos/beer-shop/project/app/user-service/api/user-service/v1"
	"github.com/go-kratos/beer-shop/project/app/user-service/internal/biz"
)

func (s *UserService) CreateCard(ctx context.Context, req *v1.CreateCardReq) (*v1.CreateCardReply, error) {
	rv, err := s.cc.Create(ctx, &biz.Card{
		CardNo:  req.CardNo,
		CCV:     req.Ccv,
		Expires: req.Expires,
	})
	return &v1.CreateCardReply{
		Id: rv.Id,
	}, err
}

func (s *UserService) GetCard(ctx context.Context, req *v1.GetCardReq) (*v1.GetCardReply, error) {
	rv, err := s.cc.Get(ctx, req.Id)
	return &v1.GetCardReply{
		Id:      rv.Id,
		CardNo:  rv.CardNo,
		Ccv:     rv.CCV,
		Expires: rv.Expires,
	}, err
}

func (s *UserService) ListCard(ctx context.Context, req *v1.ListCardReq) (*v1.ListCardReply, error) {
	rv, err := s.cc.List(ctx, req.Uid)
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
