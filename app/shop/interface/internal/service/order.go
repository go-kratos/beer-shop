package service

import (
	"context"
	"github.com/go-kratos/beer-shop/api/shop/interface/v1"
)

func (s *ShopInterface) ListCartItem(ctx context.Context, req *v1.ListCartItemReq) (*v1.ListCartItemReply, error) {
	return nil, nil
}
func (s *ShopInterface) AddCartItem(ctx context.Context, req *v1.AddCartItemReq) (*v1.AddCartItemReply, error) {
	return nil, nil
}
func (s *ShopInterface) CreateOrder(ctx context.Context, req *v1.CreateOrderReq) (*v1.CreateOrderReply, error) {
	return nil, nil
}
func (s *ShopInterface) ListOrder(ctx context.Context, req *v1.ListOrderReq) (*v1.ListOrderReply, error) {
	return nil, nil
}
