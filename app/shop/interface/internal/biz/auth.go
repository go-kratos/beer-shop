package biz

import (
	"github.com/go-kratos/beer-shop/app/shop/interface/internal/conf"
	"github.com/golang-jwt/jwt"
)

type AuthUseCase struct {
	key string
}

func NewAuthUseCase(conf *conf.Auth) *AuthUseCase {
	return &AuthUseCase{
		key: conf.ApiKey,
	}
}

func (receiver AuthUseCase) Auth(userId int64) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
	})
	return claims.SignedString(receiver.key)
}
