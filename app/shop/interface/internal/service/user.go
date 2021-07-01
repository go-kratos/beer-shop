package service

import (
	"context"
	"github.com/go-kratos/beer-shop/api/shop/interface/v1"
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
	return nil, nil
}
func (s *ShopInterface) CreateAddress(ctx context.Context, req *v1.CreateAddressReq) (*v1.CreateAddressReply, error) {
	return nil, nil
}
func (s *ShopInterface) GetAddress(ctx context.Context, req *v1.GetAddressReq) (*v1.GetAddressReply, error) {
	return nil, nil
}
func (s *ShopInterface) ListCard(ctx context.Context, req *v1.ListCardReq) (*v1.ListCardReply, error) {
	return nil, nil
}
func (s *ShopInterface) CreateCard(ctx context.Context, req *v1.CreateCardReq) (*v1.CreateCardReply, error) {
	return nil, nil
}
func (s *ShopInterface) GetCard(ctx context.Context, req *v1.GetCardReq) (*v1.GetCardReply, error) {
	return nil, nil
}
func (s *ShopInterface) DeleteCard(ctx context.Context, req *v1.DeleteCardReq) (*v1.DeleteCardReply, error) {
	return nil, nil
}
