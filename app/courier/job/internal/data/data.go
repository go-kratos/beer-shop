package data

import (
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/go-kratos/beer-shop/app/courier/job/internal/conf"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewKafkaConsumenr, NewCourierRepo)

// Data .
type Data struct {
	kc  sarama.Consumer
	log *log.Helper
}

// NewData .
func NewData(consumer sarama.Consumer, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "courier-job/data"))
	d := &Data{
		kc:  consumer,
		log: log,
	}
	return d, func() {
		d.kc.Close()
	}, nil
}

func NewKafkaConsumenr(conf *conf.Data) sarama.Consumer {
	c := sarama.NewConfig()
	p, err := sarama.NewConsumer(conf.Kafka.Addrs, c)
	if err != nil {
		panic(err)
	}
	return p
}
