package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Address struct {
	Id       int64
	Name     string
	Mobile   string
	Address  string
	PostCode string
}

type AddressRepo interface {
	CreateAddress(ctx context.Context, u *Address) (*Address, error)
	GetAddress(ctx context.Context, id int64) (*Address, error)
	ListAddress(ctx context.Context, uid int64) ([]*Address, error)
}

type AddressUseCase struct {
	repo AddressRepo
	log  *log.Helper
}

func NewAddressUseCase(repo AddressRepo, logger log.Logger) *AddressUseCase {
	return &AddressUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/address"))}
}

func (uc *AddressUseCase) Create(ctx context.Context, uid int64, a *Address) (*Address, error) {
	return uc.repo.CreateAddress(ctx, a)
}

func (uc *AddressUseCase) Get(ctx context.Context, id int64) (*Address, error) {
	return uc.repo.GetAddress(ctx, id)
}

func (uc *AddressUseCase) List(ctx context.Context, uid int64) ([]*Address, error) {
	return uc.repo.ListAddress(ctx, uid)
}
