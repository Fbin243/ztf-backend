package repo

import (
	"log"

	biz "ztf-backend/order/internal/business"
	"ztf-backend/order/internal/entity"
	"ztf-backend/pkg/db/base"

	"gorm.io/gorm"
)

type OrderRepo struct {
	*base.BaseRepo[entity.Order]
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	err := db.AutoMigrate(&entity.Order{})
	if err != nil {
		log.Printf("Error migrating order table: %v", err)
	}

	return &OrderRepo{base.NewBaseRepo[entity.Order](db)}
}

func (r *OrderRepo) WithTx(tx *gorm.DB) biz.IOrderRepo {
	return NewOrderRepo(tx)
}
