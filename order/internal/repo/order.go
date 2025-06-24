package repo

import (
	"log"

	"ztf-backend/order/internal/entity"
	"ztf-backend/shared/pkg/db"
	"ztf-backend/shared/pkg/db/base"
)

type OrderRepo struct {
	*base.BaseRepo[entity.Order]
}

func NewOrderRepo() *OrderRepo {
	err := db.GetDB().AutoMigrate(&entity.Order{})
	if err != nil {
		log.Printf("Error migrating order table: %v", err)
	}

	return &OrderRepo{base.NewBaseRepo[entity.Order]()}
}
