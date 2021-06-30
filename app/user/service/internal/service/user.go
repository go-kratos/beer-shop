package service

import (
	"context"

	v1 "github.com/go-kratos/beer-shop/api/user/service/v1"
	"github.com/go-kratos/beer-shop/app/user/service/internal/biz"
)

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserReply, error) {
	rv, err := s.uc.Create(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserReply{
		Id:       rv.Id,
		Username: rv.Username,
	}, err
}

func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	rv, err := s.uc.Get(ctx, req.Id)
	return &v1.GetUserReply{
		Id:       rv.Id,
		Username: rv.Username,
	}, err
}

func (s *UserService) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordReq) (*v1.VerifyPasswordReply, error) {
	rv, err := s.uc.VerifyPassword(ctx, &biz.User{Username: req.Username, Password: req.Password})
	return &v1.VerifyPasswordReply{
		Ok: rv,
	}, err
}
