package rpc

import (
	"context"

	biz "ztf-backend/order/internal/business"
	"ztf-backend/proto/pb/order"
)

type OrderHandler struct {
	order.UnimplementedOrderServiceServer
	userBusiness *biz.UserBusiness
}

func NewOrderHandler(userBusiness *biz.UserBusiness) *OrderHandler {
	return &OrderHandler{
		userBusiness: userBusiness,
	}
}

func (h *OrderHandler) ValidateUser(ctx context.Context, req *order.ValidateUserRequest) (*order.ValidateUserResponse, error) {
	valid, err := h.userBusiness.ValidateUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &order.ValidateUserResponse{Valid: valid}, nil
}
