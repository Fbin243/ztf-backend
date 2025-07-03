package tidb

import (
	"gorm.io/gorm"
)

type MerchantRepo struct {
	*gorm.DB
}

func NewMerchantRepo(db *gorm.DB) *MerchantRepo {
	return &MerchantRepo{db}
}
