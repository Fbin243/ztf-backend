package biz

import (
	"ztf-backend/order/internal/entity"
)

type MerchantBusiness struct {
	merchantRepo IMerchantRepo
}

func NewMerchantBusiness(merchantRepo IMerchantRepo) *MerchantBusiness {
	return &MerchantBusiness{merchantRepo: merchantRepo}
}

func (b *MerchantBusiness) FindByIds(ids []string) ([]entity.Merchant, error) {
	return b.merchantRepo.FindByIds(ids)
}
