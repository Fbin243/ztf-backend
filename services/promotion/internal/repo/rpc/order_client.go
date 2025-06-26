package rpc

import (
	"context"

	"ztf-backend/proto/pb/order"
	"ztf-backend/services/promotion/internal/entity"
)

type OrderClient struct {
	order.OrderServiceClient
}

func NewOrderClient(orderClient order.OrderServiceClient) *OrderClient {
	return &OrderClient{
		orderClient,
	}
}

func (c *OrderClient) ValidateUser(ctx context.Context, req *entity.ValidateUserReq) (bool, error) {
	response, err := c.OrderServiceClient.ValidateUser(ctx, &order.ValidateUserRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return false, err
	}

	return response.Valid, nil
}
