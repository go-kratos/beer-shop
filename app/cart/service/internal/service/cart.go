package service

import (
	"context"

	v1 "github.com/go-kratos/beer-shop/api/cart/service/v1"
	"github.com/go-kratos/beer-shop/app/cart/service/internal/biz"
)

func (s *CartService) GetCart(ctx context.Context, req *v1.GetCartReq) (reply *v1.GetCartReply, err error) {
	reply = &v1.GetCartReply{Items: make([]*v1.GetCartReply_Item, 0)}
	c, err := s.cc.GetCart(ctx, req.UserId)
	if err != nil {
		//fixme convert to error msg
		return reply, err
	}
	for _, x := range c.Items {
		reply.Items = append(reply.Items,
			&v1.GetCartReply_Item{
				ItemId:   x.Id,
				Quantity: x.Quantity,
			})
	}
	return reply, err
}
func (s *CartService) DeleteCart(ctx context.Context, req *v1.DeleteCartReq) (reply *v1.DeleteCartReply, err error) {
	reply = &v1.DeleteCartReply{}
	err = s.cc.DeleteCart(ctx, req.UserId)
	return reply, err
}
func (s *CartService) AddItem(ctx context.Context, req *v1.AddItemReq) (reply *v1.AddItemReply, err error) {
	reply = &v1.AddItemReply{}
	c, err := s.cc.AddItem(ctx, req.UserId, biz.Item{Id: req.ItemId, Quantity: req.Quantity})
	if err != nil {
		//fixme convert to error msg
		return reply, err
	}
	for _, x := range c.Items {
		reply.Items = append(reply.Items,
			&v1.AddItemReply_Item{
				ItemId:   x.Id,
				Quantity: x.Quantity,
			})
	}
	return reply, err
	return
}
func (s *CartService) UpdateItem(ctx context.Context, req *v1.UpdateItemReq) (reply *v1.UpdateItemReply, err error) {
	reply = &v1.UpdateItemReply{}
	c, err := s.cc.UpdateItem(ctx, req.UserId, req.ItemId, req.Quantity)
	if err != nil {
		//fixme convert to error msg
		return reply, err
	}
	for _, x := range c.Items {
		reply.Items = append(reply.Items,
			&v1.UpdateItemReply_Item{
				ItemId:   x.Id,
				Quantity: x.Quantity,
			})
	}
	return reply, err
	return
}
func (s *CartService) DeleteItem(ctx context.Context, req *v1.DeleteItemReq) (reply *v1.DeleteItemReply, err error) {
	reply = &v1.DeleteItemReply{}
	c, err := s.cc.DeleteItem(ctx, req.UserId, req.ItemId)
	if err != nil {
		//fixme convert to error msg
		return reply, err
	}
	for _, x := range c.Items {
		reply.Items = append(reply.Items,
			&v1.DeleteItemReply_Item{
				ItemId:   x.Id,
				Quantity: x.Quantity,
			})
	}
	return reply, err
}
