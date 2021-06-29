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

func (receiver AuthUseCase) Auth(token, username string) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token":    token,
		"username": username,
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
		result["token"] = claims["token"]
		result["username"] = claims["username"]
		return result, nil
	} else {
		return nil, errors.New("token type error")
	}
}
