package service

import (
	"context"
	v1 "github.com/go-kratos/beer-shop/api/shipping/service/v1"
	"github.com/go-kratos/beer-shop/app/shipping/service/internal/biz"
)

func (s *ShippingService) ShipOrder(ctx context.Context, req *v1.ShipOrderReq) (*v1.ShipOrderReply, error) {
	if err := s.oc.ShipOrder(ctx, &biz.ShipOrder{Id: req.Id}); err != nil {
		return nil, err
	}

	return &v1.ShipOrderReply{
		Id: req.Id,
	}, nil
}
