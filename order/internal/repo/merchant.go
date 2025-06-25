package repo

import (
	"ztf-backend/order/internal/entity"
	"ztf-backend/pkg/db/base"

	"gorm.io/gorm"
)

type MerchantRepo struct {
	*base.BaseRepo[entity.Merchant]
}

func NewMerchantRepo(db *gorm.DB) *MerchantRepo {
	return &MerchantRepo{base.NewBaseRepo[entity.Merchant](db)}
}
