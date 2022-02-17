package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/beer-shop/app/user/service/internal/biz"
	"github.com/go-kratos/beer-shop/app/user/service/internal/data/ent"
	"github.com/go-kratos/beer-shop/app/user/service/internal/data/ent/user"
	"github.com/go-kratos/beer-shop/app/user/service/internal/pkg/util"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*userRepo)(nil)

var userCacheKey = func(username string) string {
	return "user_cache_key_" + username
}

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

func (r *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	var target *ent.User
	// try to fetch from cache
	cacheKey := userCacheKey(username)
	target, err := r.getUserFromCache(ctx, cacheKey)
	if err != nil {
		// fetch from db while cache missed
		target, err = r.data.db.User.
			Query().
			Where(user.UsernameEQ(username)).
			Only(ctx)
		if err != nil {
			return nil, biz.ErrUserNotFound
		}
		// set cache
		r.serUserCache(ctx, target, cacheKey)
	}
	return &biz.User{
		Id:       target.ID,
		Username: target.Username,
		Password: target.PasswordHash,
	}, nil
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	ph, err := util.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	po, err := r.data.db.User.
		Create().
		SetUsername(u.Username).
		SetPasswordHash(ph).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{Id: po.ID, Username: po.Username}, nil
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	// try to fetch from cache
	cacheKey := userCacheKey(fmt.Sprintf("%d", id))
	target, err := r.getUserFromCache(ctx, cacheKey)
	if err != nil {
		// fetch from db while cache missed
		target, err = r.data.db.User.Get(ctx, id)
		if err != nil {
			return nil, biz.ErrUserNotFound
		}
		// set cache
		r.serUserCache(ctx, target, cacheKey)
	}
	return &biz.User{Id: target.ID, Username: target.Username}, nil
}

func (r *userRepo) VerifyPassword(ctx context.Context, u *biz.User) (bool, error) {
	po, err := r.data.db.User.
		Query().
		Where(user.UsernameEQ(u.Username)).
		Only(ctx)
	if err != nil {
		return false, err
	}
	return util.CheckPasswordHash(u.Password, po.PasswordHash), nil
}

func (r *userRepo) getUserFromCache(ctx context.Context, key string) (*ent.User, error) {
	result, err := r.data.redisCli.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var cacheUser = &ent.User{}
	err = json.Unmarshal([]byte(result), cacheUser)
	if err != nil {
		return nil, err
	}
	return cacheUser, nil
}

func (r *userRepo) serUserCache(ctx context.Context, user *ent.User, key string) {
	marshal, err := json.Marshal(user)
	if err != nil {
		r.log.Errorf("fail to set user cache:json.Marshal(%v) error(%v)", user, err)
	}
	err = r.data.redisCli.Set(ctx, key, string(marshal), time.Minute*30).Err()
	if err != nil {
		r.log.Errorf("fail to set user cache:redis.Set(%v) error(%v)", user, err)
	}
}
