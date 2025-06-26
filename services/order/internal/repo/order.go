package repo

import (
	"log"

	"ztf-backend/pkg/db/base"
	biz "ztf-backend/services/order/internal/business"
	"ztf-backend/services/order/internal/entity"

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
