package server

import (
	"github.com/go-kratos/beer-shop/api/user/service/v1"
	"github.com/go-kratos/beer-shop/app/user/service/internal/conf"
	"github.com/go-kratos/beer-shop/app/user/service/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwt2 "github.com/golang-jwt/jwt/v4"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, ac *conf.Auth, logger log.Logger, tp *tracesdk.TracerProvider, s *service.UserService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(
				tracing.WithTracerProvider(tp)),
			logging.Server(logger),
			jwt.Server(func(token *jwt2.Token) (interface{}, error) {
				return []byte(ac.Key), nil
			}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterUserServer(srv, s)
	return srv
}
