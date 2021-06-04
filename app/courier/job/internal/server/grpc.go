package server

import (
	v1 "github.com/go-kratos/beer-shop/api/courier/job/v1"
	"github.com/go-kratos/beer-shop/app/courier/job/internal/conf"
	"github.com/go-kratos/beer-shop/app/courier/job/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/propagation"

	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, logger log.Logger, tp *tracesdk.TracerProvider, s *service.CourierService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(
				tracing.WithTracerProvider(tp),
				tracing.WithPropagators(
					propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}),
				),
			),
			logging.Server(logger),
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
	v1.RegisterCourierServer(srv, s)
	return srv
}
