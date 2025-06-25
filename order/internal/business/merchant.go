package biz

import (
	"context"

	"ztf-backend/order/internal/entity"
)

type MerchantBusiness struct {
	merchantRepo IMerchantRepo
}

func NewMerchantBusiness(merchantRepo IMerchantRepo) *MerchantBusiness {
	return &MerchantBusiness{merchantRepo: merchantRepo}
}

func (b *MerchantBusiness) FindByIds(ctx context.Context, ids []string) ([]entity.Merchant, error) {
	return b.merchantRepo.FindByIds(ctx, ids)
}
