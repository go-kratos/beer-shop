package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go.opentelemetry.io/otel/propagation"

	"github.com/go-kratos/beer-shop/app/shop/interface/internal/conf"

	"context"

	csV1 "github.com/go-kratos/beer-shop/api/cart/service/v1"
	ctV1 "github.com/go-kratos/beer-shop/api/catalog/service/v1"
	osV1 "github.com/go-kratos/beer-shop/api/order/service/v1"
	psV1 "github.com/go-kratos/beer-shop/api/payment/service/v1"
	usV1 "github.com/go-kratos/beer-shop/api/user/service/v1"

	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDiscovery,
	NewUserServiceClient,
	NewCartServiceClient,
	NewCatalogServiceClient,
	NewOrderServiceClient,
	NewPaymentServiceClient,
	NewUserRepo,
)

// Data .
type Data struct {
	log *log.Helper
	uc  usV1.UserClient
	cc  csV1.CartClient
	bc  ctV1.CatalogClient
}

// NewData .
func NewData(
	conf *conf.Data,
	logger log.Logger,
	uc usV1.UserClient,
	cc csV1.CartClient,
	bc ctV1.CatalogClient,
) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{log: l, uc: uc, cc: cc, bc: bc}, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli)
	return r
}

func NewUserServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider) usV1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///beer.user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(middleware.Chain(
			tracing.Client(
				tracing.WithTracerProvider(tp),
				tracing.WithPropagators(
					propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}),
				),
			),
			recovery.Recovery())),
	)
	if err != nil {
		panic(err)
	}
	c := usV1.NewUserClient(conn)
	return c
}

func NewCartServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider) csV1.CartClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///beer.cart.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(middleware.Chain(
			tracing.Client(
				tracing.WithTracerProvider(tp),
				tracing.WithPropagators(
					propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}),
				),
			),
			recovery.Recovery())),
	)
	if err != nil {
		panic(err)
	}
	return csV1.NewCartClient(conn)
}

func NewCatalogServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider) ctV1.CatalogClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///beer.catalog.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(middleware.Chain(
			tracing.Client(
				tracing.WithTracerProvider(tp),
				tracing.WithPropagators(
					propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}),
				),
			),
			recovery.Recovery())),
	)
	if err != nil {
		panic(err)
	}
	return ctV1.NewCatalogClient(conn)
}

func NewOrderServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider) osV1.OrderClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///beer.order.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(middleware.Chain(
			tracing.Client(
				tracing.WithTracerProvider(tp),
				tracing.WithPropagators(
					propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}),
				),
			),
			recovery.Recovery())),
	)
	if err != nil {
		panic(err)
	}
	return osV1.NewOrderClient(conn)
}

func NewPaymentServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider) psV1.PaymentClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///beer.payment.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(middleware.Chain(
			tracing.Client(
				tracing.WithTracerProvider(tp),
				tracing.WithPropagators(
					propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}),
				),
			),
			recovery.Recovery())),
	)
	if err != nil {
		panic(err)
	}
	return psV1.NewPaymentClient(conn)
}
