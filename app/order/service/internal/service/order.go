package service

import (
	"context"

	v1 "github.com/go-kratos/beer-shop/app/order/service/internal/api/order/service/v1"
	"github.com/go-kratos/beer-shop/app/order/service/internal/biz"
)

func (s *OrderService) CreateOrder(ctx context.Context, req *v1.CreateOrderReq) (*v1.CreateOrderReply, error) {
	x, err := s.oc.Create(ctx, &biz.Order{})
	return &v1.CreateOrderReply{
		Id: x.Id,
	}, err
}

func (s *OrderService) GetOrder(ctx context.Context, req *v1.GetOrderReq) (*v1.GetOrderReply, error) {
	x, err := s.oc.Get(ctx, req.Id)
	return &v1.GetOrderReply{
		Id: x.Id,
	}, err
}

func (s *OrderService) UpdateOrder(ctx context.Context, req *v1.UpdateOrderReq) (*v1.UpdateOrderReply, error) {
	x, err := s.oc.Update(ctx, &biz.Order{})
	return &v1.UpdateOrderReply{
		Id: x.Id,
	}, err
}

func (s *OrderService) ListOrder(ctx context.Context, req *v1.ListOrderReq) (*v1.ListOrderReply, error) {
	rv, err := s.oc.List(ctx, req.PageNum, req.PageSize)
	rs := make([]*v1.ListOrderReply_Order, 0)
	for _, x := range rv {
		rs = append(rs, &v1.ListOrderReply_Order{
			Id: x.Id,
		})
	}
	return &v1.ListOrderReply{
		Orders: rs,
	}, err
}
