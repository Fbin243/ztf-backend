package repo

import (
	"ztf-backend/internal/db"
	"ztf-backend/internal/entity"
)

type MerchantRepo struct {
	*db.BaseRepo[entity.Merchant]
}

func NewMerchantRepo() *MerchantRepo {
	return &MerchantRepo{db.NewBaseRepo[entity.Merchant]()}
}
