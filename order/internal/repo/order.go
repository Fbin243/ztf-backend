package repo

import (
	"ztf-backend/order/internal/entity"
	"ztf-backend/shared/pkg/db/base"
)

type OrderRepo struct {
	*base.BaseRepo[entity.Order]
}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{base.NewBaseRepo[entity.Order]()}
}
