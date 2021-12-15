package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewData, NewPaymentRepo)

// Data .
type Data struct {
	log *log.Helper
}

// NewData .
func NewData(logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "payment-service/data"))

	d := &Data{
		log: log,
	}
	return d, func() {
	}, nil
}
