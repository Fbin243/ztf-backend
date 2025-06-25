package biz

import (
	"context"

	"ztf-backend/promotion/internal/entity"
)

func (b *PromotionBusiness) FindAll(ctx context.Context) ([]entity.Promotion, error) {
	return b.promotionRepo.FindAll(ctx)
}

func (b *PromotionBusiness) FindById(ctx context.Context, id string) (*entity.Promotion, error) {
	return b.promotionRepo.FindById(ctx, id)
}
