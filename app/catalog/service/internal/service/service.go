package service

import (
	v1 "github.com/go-kratos/beer-shop/app/catalog/service/internal/api/catalog/service/v1"
	"github.com/go-kratos/beer-shop/app/catalog/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCatalogService)

type CatalogService struct {
	v1.UnimplementedCatalogServer

	bc  *biz.BeerUseCase
	log *log.Helper
}

func NewCatalogService(bc *biz.BeerUseCase, logger log.Logger) *CatalogService {
	return &CatalogService{

		bc:  bc,
		log: log.NewHelper(log.With(logger, "module", "service/catalog"))}
}
