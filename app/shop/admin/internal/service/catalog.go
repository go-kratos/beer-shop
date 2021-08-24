package service

import (
	"context"
	"github.com/go-kratos/beer-shop/api/shop/admin/v1"
)

func (s *ShopAdmin) ListBeer(ctx context.Context, req *v1.ListBeerReq) (*v1.ListBeerReply, error) {
	rv, err := s.cc.ListBeer(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	reply := &v1.ListBeerReply{
		Results: make([]*v1.ListBeerReply_Beer, 0),
	}
	for _, x := range rv {
		item := &v1.ListBeerReply_Beer{
			Id:          x.Id,
			Name:        x.Name,
			Description: x.Description,
			Count:       x.Count,
			Image:       make([]*v1.ListBeerReply_Beer_Image, 0),
		}
		for _, img := range x.Images {
			item.Image = append(item.Image, &v1.ListBeerReply_Beer_Image{Url: img.URL})
		}
		reply.Results = append(reply.Results, item)
	}
	return reply, nil
}
func (s *ShopAdmin) GetBeer(ctx context.Context, req *v1.GetBeerReq) (*v1.GetBeerReply, error) {
	x, err := s.cc.GetBeer(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	reply := &v1.GetBeerReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       make([]*v1.GetBeerReply_Image, 0),
	}
	for _, img := range x.Images {
		reply.Image = append(reply.Image, &v1.GetBeerReply_Image{Url: img.URL})
	}
	return reply, nil
}

func (s *ShopAdmin) CreateBeer(ctx context.Context, req *v1.CreateBeerReq) (*v1.CreateBeerReply, error) {
	return nil, nil
}

func (s *ShopAdmin) UpdateBeer(ctx context.Context, req *v1.UpdateBeerReq) (*v1.UpdateBeerReply, error) {
	return nil, nil
}

func (s *ShopAdmin) DeleteBeer(ctx context.Context, req *v1.DeleteBeerReq) (*v1.DeleteBeerReply, error) {
	return nil, nil
}
