package service

import (
	"context"
	"github.com/go-kratos/beer-shop/api/cart/service/v1"
)

func (s *CartService) GetCart(ctx context.Context, req *v1.GetCartReq) (reply *v1.GetCartReq, err error) {
	return
}
func (s *CartService) DeleteCart(ctx context.Context, req *v1.DeleteCartReq) (reply *v1.DeleteCartReply, err error) {
	return
}
func (s *CartService) AddItem(ctx context.Context, req *v1.AddItemReq) (reply *v1.AddItemReply, err error) {
	return
}
func (s *CartService) UpdateItem(ctx context.Context, req *v1.UpdateItemReq) (reply *v1.UpdateItemReply, err error) {
	return
}
func (s *CartService) DeleteItem(ctx context.Context, req *v1.DeleteItemReq) (reply *v1.DeleteItemReply, err error) {
	return
}
