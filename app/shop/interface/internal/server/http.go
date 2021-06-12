package server

import (
	"github.com/go-kratos/beer-shop/api/shop/interface/v1"
	"github.com/go-kratos/beer-shop/app/shop/interface/internal/conf"
	"github.com/go-kratos/beer-shop/app/shop/interface/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/propagation"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, tp *tracesdk.TracerProvider, s *service.ShopInterface) *http.Server {
	var opts = []http.ServerOption{}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	m := http.Middleware(
		recovery.Recovery(),
		tracing.Server(
			tracing.WithTracerProvider(tp),
			tracing.WithPropagators(
				propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}),
			),
		),
		logging.Server(logger),
	)
	srv.HandlePrefix("/", v1.NewShopInterfaceHandler(s, m))
	return srv
}
