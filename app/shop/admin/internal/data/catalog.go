package data

import (
	"context"

	"github.com/go-kratos/beer-shop/app/shop/admin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	catalogv1 "github.com/go-kratos/beer-shop/api/catalog/service/v1"
)

var _ biz.CatalogRepo = (*catalogRepo)(nil)

type catalogRepo struct {
	data *Data
	log  *log.Helper
}

func NewCatalogRepo(data *Data, logger log.Logger) biz.CatalogRepo {
	return &catalogRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/beer")),
	}
}

func (r *catalogRepo) GetBeer(ctx context.Context, id int64) (*biz.Beer, error) {
	reply, err := r.data.bc.GetBeer(ctx, &catalogv1.GetBeerReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	images := make([]biz.Image, 0)
	for _, x := range reply.Image {
		images = append(images, biz.Image{URL: x.Url})
	}
	return &biz.Beer{
		Id:          reply.Id,
		Name:        reply.Name,
		Description: reply.Description,
		Count:       reply.Count,
		Images:      images,
	}, err
}

func (r *catalogRepo) ListBeer(ctx context.Context, pageNum, pageSize int64) ([]*biz.Beer, error) {
	reply, err := r.data.bc.ListBeer(ctx, &catalogv1.ListBeerReq{
		PageNum:  pageNum,
		PageSize: pageSize,
	})
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Beer, 0)
	for _, x := range reply.Results {
		images := make([]biz.Image, 0)
		for _, img := range x.Image {
			images = append(images, biz.Image{URL: img.Url})
		}
		rv = append(rv, &biz.Beer{
			Id:          x.Id,
			Description: x.Description,
			Count:       x.Count,
			Images:      images,
		})
	}
	return rv, err
}

func (r *catalogRepo) CreateBeer(ctx context.Context, b *biz.Beer) (*biz.Beer, error) {
	images := make([]*catalogv1.CreateBeerReq_Image, 0)
	for _, x := range b.Images {
		images = append(images, &catalogv1.CreateBeerReq_Image{Url: x.URL})
	}
	reply, err := r.data.bc.CreateBeer(ctx, &catalogv1.CreateBeerReq{
		Name:        b.Name,
		Description: b.Description,
		Count:       b.Count,
		Image:       images,
	})
	if err != nil {
		return nil, err
	}

	return &biz.Beer{
		Id: reply.Id,
	}, err
}

func (r *catalogRepo) UpdateBeer(ctx context.Context, b *biz.Beer) (*biz.Beer, error) {
	images := make([]*catalogv1.UpdateBeerReq_Image, 0)
	for _, x := range b.Images {
		images = append(images, &catalogv1.UpdateBeerReq_Image{Url: x.URL})
	}
	reply, err := r.data.bc.UpdateBeer(ctx, &catalogv1.UpdateBeerReq{
		Name:        b.Name,
		Description: b.Description,
		Count:       b.Count,
		Image:       images,
	})
	if err != nil {
		return nil, err
	}

	return &biz.Beer{
		Id: reply.Id,
	}, err
}
