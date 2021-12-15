package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/beer-shop/app/shipping/service/internal/biz"
)

var _ biz.ShippingRepo = (*shippingRepo)(nil)

type shippingRepo struct {
	data *Data
	log  *log.Helper
}

func NewShippingRepo(data *Data, logger log.Logger) biz.ShippingRepo {
	return &shippingRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/shipping")),
	}
}

type ShippingEntry struct {
	OrderId string `json:"order_id"`
}

func (uc *shippingRepo) ShipOrder(ctx context.Context, o *biz.ShipOrder) (err error) {
	se := &ShippingEntry{
		OrderId: fmt.Sprintf("%d", o.Id),
	}
	b, err := json.Marshal(se)
	if err != nil {
		return err
	}
	uc.data.kp.Input() <- &sarama.ProducerMessage{
		Topic: "shipping",
		Value: sarama.ByteEncoder(b),
	}
	return
}
