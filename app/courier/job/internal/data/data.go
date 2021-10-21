package data

import (
	"context"
	"github.com/Shopify/sarama"
	orderv1 "github.com/go-kratos/beer-shop/api/order/service/v1"
	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-kratos/beer-shop/app/courier/job/internal/conf"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewKafkaConsumer,
	NewCourierRepo,
	NewDiscovery,
	NewOrderServiceClient,
)

// Data .
type Data struct {
	kc  sarama.Consumer
	oc  orderv1.OrderClient
	log *log.Helper
}

// NewData .
func NewData(consumer sarama.Consumer, logger log.Logger, oc orderv1.OrderClient,
) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "courier-job/data"))
	d := &Data{
		kc:  consumer,
		oc:  oc,
		log: log,
	}
	return d, func() {
		d.kc.Close()
	}, nil
}

func NewKafkaConsumer(conf *conf.Data) sarama.Consumer {
	c := sarama.NewConfig()
	p, err := sarama.NewConsumer(conf.Kafka.Addrs, c)
	if err != nil {
		panic(err)
	}
	return p
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewOrderServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider) orderv1.OrderClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///beer.order.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return orderv1.NewOrderClient(conn)
}
