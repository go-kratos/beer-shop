package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/go-kratos/beer-shop/app/shop/admin/internal/conf"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	log *log.Helper
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "server-service/data"))
	return &Data{log: l}, nil
}
