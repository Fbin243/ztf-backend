package repo

import (
	"ztf-backend/pkg/db/base"
	"ztf-backend/services/order/internal/entity"

	"gorm.io/gorm"
)

type MerchantRepo struct {
	*base.BaseRepo[entity.Merchant]
}

func NewMerchantRepo(db *gorm.DB) *MerchantRepo {
	return &MerchantRepo{base.NewBaseRepo[entity.Merchant](db)}
}
