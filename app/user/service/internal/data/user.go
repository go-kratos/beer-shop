package data

import (
	"context"

	"github.com/go-kratos/beer-shop/app/user/service/internal/biz"
	"github.com/go-kratos/beer-shop/app/user/service/internal/data/ent/user"
	"github.com/go-kratos/beer-shop/app/user/service/internal/pkg/utils"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	ph, err := utils.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	po, err := r.data.db.User.
		Create().
		SetUsername(u.Username).
		SetPasswordHash(ph).
		Save(ctx)
	return &biz.User{Id: po.ID, Username: po.Username}, err
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	po, err := r.data.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.User{Id: po.ID, Username: po.Username}, err
}

func (r *userRepo) VerifyPassword(ctx context.Context, u *biz.User) (bool, error) {
	po, err := r.data.db.User.
		Query().
		Where(user.UsernameEQ(u.Username)).
		Only(ctx)
	if err != nil {
		return false, err
	}
	return utils.CheckPasswordHash(u.Password, po.PasswordHash), nil
}
