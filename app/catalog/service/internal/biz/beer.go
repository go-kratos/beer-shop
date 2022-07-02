package biz

import (
	"context"
	"github.com/go-kratos/beer-shop/pkg/page_token"
	"golang.org/x/sync/singleflight"

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
	ListBeerNext(ctx context.Context, start, end int32) ([]*Beer, error)
	Count(ctx context.Context) (int, error)
}

type BeerUseCase struct {
	repo         BeerRepo
	log          *log.Helper
	pageToken    page_token.ProcessPageTokens
	singleflight singleflight.Group
}

func NewBeerUseCase(repo BeerRepo, logger log.Logger) *BeerUseCase {
	return &BeerUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/beer")), pageToken: page_token.NewTokenGenerate()}
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

func (uc *BeerUseCase) ListNext(ctx context.Context, pageSize int32, pageToken string) ([]*Beer, string, error) {
	total, err := uc.repo.Count(ctx)
	if err != nil {
		return nil, "", err
	}

	start, end, nextToken, err := uc.pageToken.ProcessPageTokens(total, pageSize, pageToken)
	if err != nil {
		return nil, "", err
	}
	// singleflight
	data, err, _ := uc.singleflight.Do("list_next", func() (interface{}, error) {
		return uc.repo.ListBeerNext(ctx, start, end)
	})

	return data.([]*Beer), nextToken, err
}
