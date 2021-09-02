package biz

import (
	"errors"
	"github.com/go-kratos/beer-shop/app/shop/interface/internal/conf"
	"github.com/golang-jwt/jwt"
)

type AuthUseCase struct {
	key string
}

func NewAuthUseCase(conf *conf.Auth) *AuthUseCase {
	return &AuthUseCase{
		key: conf.Key,
	}
}

func (receiver AuthUseCase) Auth(userId int64) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
	})
	return claims.SignedString(receiver.key)
}

func (receiver AuthUseCase) CheckJWT(jwtToken string) (map[string]interface{}, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return receiver.key, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		result := make(map[string]interface{}, 2)
		result["user_id"] = claims["user_id"]
		return result, nil
	} else {
		return nil, errors.New("token type error")
	}
}
