package repo

import (
	"ztf-backend/internal/db"
	"ztf-backend/internal/entity"
)

type OrderRepo struct {	
	*db.BaseRepo[entity.Order]
}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{db.NewBaseRepo[entity.Order]()}
}
