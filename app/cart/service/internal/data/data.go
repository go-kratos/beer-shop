package data

import (
	"context"
	"github.com/go-kratos/beer-shop/app/cart/service/internal/data/ent"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"

	"github.com/go-kratos/beer-shop/app/cart/service/internal/conf"

)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCartRepo)

// Data .
type Data struct {
	db *mongo.Client
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, error) {
	log := log.NewHelper("cart-service/data", logger)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	err = client.Ping(ctx, readpref.Primary())


	return &Data{
		db: client,
	}, nil
}
