package tidb

import (
	"log/slog"
	"ztf-backend/services/order/internal/entity"

	biz "ztf-backend/services/order/internal/business"

	"gorm.io/gorm"
)

type OrderRepo struct {
	*gorm.DB
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	err := db.AutoMigrate(&entity.Order{})
	if err != nil {
		slog.Error(err.Error())
	}

	return &OrderRepo{db}
}

func (r *OrderRepo) WithTx(tx *gorm.DB) biz.IOrderRepo {
	return NewOrderRepo(tx)
}
