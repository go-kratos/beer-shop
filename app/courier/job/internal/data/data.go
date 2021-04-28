package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/go-kratos/beer-shop/app/courier/job/internal/conf"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCourierRepo)

// Data .
type Data struct {
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	d := &Data{}
	return d, func() {

	}, nil
}
