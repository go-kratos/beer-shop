package cors

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"

	"github.com/go-kratos/kratos/v2/middleware"
	http2 "github.com/go-kratos/kratos/v2/transport/http"
)

func CORS() middleware.Middleware {
	logger := log.NewHelper(log.DefaultLogger)
	logger.Info("CORS registered")
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if info, ok := http2.FromServerContext(ctx); ok {
				logger.Info("cors request")
				info.Response.Header().Set("Access-Control-Allow-Origin", "*")
				if info.Request.Method == http.MethodOptions {
					return
				}
			}
			return handler(ctx, req)
		}
	}
}
