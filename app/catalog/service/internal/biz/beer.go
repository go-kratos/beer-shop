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

type BeerRepo interface {
	CreateBeer(ctx context.Context, c *Beer) (*Beer, error)
	UpdateBeer(ctx context.Context, c *Beer) (*Beer, error)
	GetBeer(ctx context.Context, id int64) (*Beer, error)
	ListBeer(ctx context.Context, pageNum, pageSize int64) ([]*Beer, error)
}

type BeerUseCase struct {
	repo BeerRepo
	log  *log.Helper
}

func NewBeerUseCase(repo BeerRepo, logger log.Logger) *BeerUseCase {
	return &BeerUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/beer"))}
}

func (uc *BeerUseCase) Create(ctx context.Context, u *Beer) (*Beer, error) {
	return uc.repo.CreateBeer(ctx, u)
}

func (uc *BeerUseCase) Get(ctx context.Context, id int64) (*Beer, error) {
	return uc.repo.GetBeer(ctx, id)
}

func (uc *BeerUseCase) Update(ctx context.Context, u *Beer) (*Beer, error) {
	return uc.repo.UpdateBeer(ctx, u)
}

func (uc *BeerUseCase) List(ctx context.Context, pageNum, pageSize int64) ([]*Beer, error) {
	return uc.repo.ListBeer(ctx, pageNum, pageSize)
}
