package jwt

import (
	"context"
	"github.com/go-kratos/beer-shop/app/shop/interface/internal/biz"
	"github.com/go-kratos/kratos/v2/middleware"
	"net/http"
)

func NewAuthMiddleware(authUc *biz.AuthUseCase) func(handler middleware.Handler) middleware.Handler {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if request, ok := req.(http.Request); ok {
				jwtToken := request.Header.Get("Authorization")
				token, err := authUc.CheckJWT(jwtToken)
				if err != nil {
					// todo 这里不知道怎么处理reply
					return nil, err
				}
				ctx = context.WithValue(ctx, "x-md-global-uid", token["user_id"])
			}
			reply, err = handler(ctx, req)
			return
		}
	}

}
