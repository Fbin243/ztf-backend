package transport

import biz "ztf-backend/internal/business"

type OrderHandler struct {
	orderBusiness *biz.OrderBusiness
}

func NewOrderHandler(orderBusiness *biz.OrderBusiness) *OrderHandler {
	return &OrderHandler{
		orderBusiness: orderBusiness,
	}
}
