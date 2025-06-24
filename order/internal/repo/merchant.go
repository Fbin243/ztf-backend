package repo

import (
	"ztf-backend/order/internal/entity"
	"ztf-backend/shared/pkg/db/base"
)

type MerchantRepo struct {
	*base.BaseRepo[entity.Merchant]
}

func NewMerchantRepo() *MerchantRepo {
	return &MerchantRepo{base.NewBaseRepo[entity.Merchant]()}
}
