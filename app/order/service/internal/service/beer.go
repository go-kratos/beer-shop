package service

import (
	"context"
	"github.com/go-kratos/beer-shop/api/order/service/v1"
	"github.com/go-kratos/beer-shop/app/order/service/internal/biz"
)

func (s *OrderService) CreateBeer(ctx context.Context, req *v1.CreateBeerReq) (*v1.CreateBeerReply, error) {
	x, err := s.bc.Create(ctx, &biz.Beer{
	})
	return &v1.CreateBeerReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
	}, err
}

func (s *OrderService) GetBeer(ctx context.Context, req *v1.GetBeerReq) (*v1.GetBeerReply, error) {
	x, err := s.bc.Get(ctx, req.Id)
	return &v1.GetBeerReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
	}, err
}

func (s *OrderService) ListBeer(ctx context.Context, req *v1.ListBeerReq) (*v1.ListBeerReply, error) {
	rv, err := s.bc.List(ctx, req.PageNum, req.PageSize)
	rs := make([]*v1.ListBeerReply_Beer, 0)
	for _, x := range rv {
		rs = append(rs, &v1.ListBeerReply_Beer{
			Id:          x.Id,
			Name:        x.Name,
			Description: x.Description,
		})
	}
	return &v1.ListBeerReply{
		Results: rs,
	}, err
}
