package tidb

import (
	"gorm.io/gorm"
	"ztf-backend/pkg/db/base"
	"ztf-backend/services/order/internal/entity"
)

type MerchantRepo struct {
	*base.BaseRepo[entity.Merchant]
}

func NewMerchantRepo(db *gorm.DB) *MerchantRepo {
	return &MerchantRepo{base.NewBaseRepo[entity.Merchant](db)}
}
