package service

import (
	"context"
	"github.com/go-kratos/beer-shop/api/shop/interface/v1"
)

func (s *ShopInterface) ListBeer(ctx context.Context, req *v1.ListBeerReq) (*v1.ListBeerReply, error) {
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
func (s *ShopInterface) GetBeer(ctx context.Context, req *v1.GetBeerReq) (*v1.GetBeerReply, error) {
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
