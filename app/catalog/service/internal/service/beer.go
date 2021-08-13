package service

import (
	"context"

	v1 "github.com/go-kratos/beer-shop/app/catalog/service/internal/api/catalog/service/v1"
	"github.com/go-kratos/beer-shop/app/catalog/service/internal/biz"
)

func (s *CatalogService) CreateBeer(ctx context.Context, req *v1.CreateBeerReq) (*v1.CreateBeerReply, error) {
	b := &biz.Beer{
		Name:        req.Name,
		Description: req.Description,
		Count:       req.Count,
		Images:      make([]biz.Image, 0),
	}
	for _, x := range req.Image {
		b.Images = append(b.Images, biz.Image{URL: x.Url})
	}
	x, err := s.bc.Create(ctx, b)
	img := make([]*v1.CreateBeerReply_Image, 0)
	for _, i := range x.Images {
		img = append(img, &v1.CreateBeerReply_Image{Url: i.URL})
	}
	return &v1.CreateBeerReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       img,
	}, err
}

func (s *CatalogService) GetBeer(ctx context.Context, req *v1.GetBeerReq) (*v1.GetBeerReply, error) {
	x, err := s.bc.Get(ctx, req.Id)
	img := make([]*v1.GetBeerReply_Image, 0)
	for _, i := range x.Images {
		img = append(img, &v1.GetBeerReply_Image{Url: i.URL})
	}
	return &v1.GetBeerReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       img,
	}, err
}

func (s *CatalogService) UpdateBeer(ctx context.Context, req *v1.UpdateBeerReq) (*v1.UpdateBeerReply, error) {
	b := &biz.Beer{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Count:       req.Count,
		Images:      make([]biz.Image, 0),
	}
	for _, x := range req.Image {
		b.Images = append(b.Images, biz.Image{URL: x.Url})
	}
	x, err := s.bc.Update(ctx, b)
	img := make([]*v1.UpdateBeerReply_Image, 0)
	for _, i := range x.Images {
		img = append(img, &v1.UpdateBeerReply_Image{Url: i.URL})
	}
	return &v1.UpdateBeerReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       img,
	}, err
}

func (s *CatalogService) ListBeer(ctx context.Context, req *v1.ListBeerReq) (*v1.ListBeerReply, error) {
	rv, err := s.bc.List(ctx, req.PageNum, req.PageSize)
	rs := make([]*v1.ListBeerReply_Beer, 0)
	for _, x := range rv {
		img := make([]*v1.ListBeerReply_Beer_Image, 0)
		for _, i := range x.Images {
			img = append(img, &v1.ListBeerReply_Beer_Image{Url: i.URL})
		}
		rs = append(rs, &v1.ListBeerReply_Beer{
			Id:          x.Id,
			Name:        x.Name,
			Description: x.Description,
			Image:       img,
		})
	}
	return &v1.ListBeerReply{
		Results: rs,
	}, err
}
