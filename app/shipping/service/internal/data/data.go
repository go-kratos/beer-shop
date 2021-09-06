package data

import (
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/go-kratos/beer-shop/app/shipping/service/internal/conf"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewKafkaProducer, NewShippingRepo)

// Data .
type Data struct {
	kp  sarama.AsyncProducer
	log *log.Helper
}

// NewData .
func NewData(producer sarama.AsyncProducer, conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "shipping-service/data"))
	d := &Data{
		kp:  producer,
		log: log,
	}
	return d, func() {
		d.kp.Close()
	}, nil
}

func NewKafkaProducer(conf *conf.Data) sarama.AsyncProducer {
	c := sarama.NewConfig()
	p, err := sarama.NewAsyncProducer(conf.Kafka.Addrs, c)
	if err != nil {
		panic(err)
	}
	return p
}
