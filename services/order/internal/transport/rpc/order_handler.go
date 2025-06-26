package rpc

import (
	"context"

	"ztf-backend/proto/pb/order"
	biz "ztf-backend/services/order/internal/business"
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

func (h *OrderHandler) ValidateUser(
	ctx context.Context,
	req *order.ValidateUserRequest,
) (*order.ValidateUserResponse, error) {
	valid, err := h.userBusiness.ValidateUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &order.ValidateUserResponse{Valid: valid}, nil
}
