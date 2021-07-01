package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Image struct {
	URL string
}

type Beer struct {
	Id          int64
	Name        string
	Description string
	Count       int64
	Images      []Image
}

type CatalogRepo interface {
	GetBeer(ctx context.Context, id int64) (*Beer, error)
	ListBeer(ctx context.Context, pageNum, pageSize int64) ([]*Beer, error)
}

type CatalogUseCase struct {
	repo CatalogRepo
	log  *log.Helper
}

func NewCatalogUseCase(repo CatalogRepo, logger log.Logger) *CatalogUseCase {
	return &CatalogUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/beer"))}
}

func (uc *CatalogUseCase) GetBeer(ctx context.Context, id int64) (*Beer, error) {
	return uc.repo.GetBeer(ctx, id)
}

func (uc *CatalogUseCase) ListBeer(ctx context.Context, pageNum, pageSize int64) ([]*Beer, error) {
	return uc.repo.ListBeer(ctx, pageNum, pageSize)
}
