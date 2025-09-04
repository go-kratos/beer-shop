package data

import (
	"context"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/XSAM/otelsql"
	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"time"

	"github.com/go-kratos/beer-shop/app/user/service/internal/conf"
	"github.com/go-kratos/beer-shop/app/user/service/internal/data/ent"
	"github.com/go-kratos/beer-shop/app/user/service/internal/data/ent/migrate"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewRedisCmd, NewUserRepo, NewCardRepo, NewAddressRepo)

// Data .
type Data struct {
	db       *ent.Client
	redisCli redis.Cmdable
}

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "user-service/data/ent"))

	// Wrap over sql.Open with OTel instrumentation.
	attributes := otelsql.WithAttributes(
		semconv.DBSystemKey.String(conf.Database.Driver),
	)
	spanOptions := otelsql.WithSpanOptions(
		otelsql.SpanOptions{
			DisableErrSkip: true,
		})
	db, err := otelsql.Open(conf.Database.Driver, conf.Database.Source, attributes, spanOptions)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}

	// New an ent client.
	drv := entsql.OpenDB(conf.Database.Driver, db)
	client := ent.NewClient(ent.Driver(drv))

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func NewRedisCmd(conf *conf.Data, logger log.Logger) redis.Cmdable {
	log := log.NewHelper(log.With(logger, "module", "user-service/data/ent"))
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})
	client.AddHook(redisotel.NewTracingHook())
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	err := client.Ping(timeout).Err()
	if err != nil {
		log.Fatalf("redis connect error: %v", err)
	}
	return client
}

// NewData .
func NewData(entClient *ent.Client, redisCmd redis.Cmdable, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "user-service/data"))

	d := &Data{
		db:       entClient,
		redisCli: redisCmd,
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
