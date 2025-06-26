package biz

import (
	"context"

	"ztf-backend/services/promotion/internal/entity"
)

func (b *PromotionBusiness) FindAll(ctx context.Context) ([]entity.Promotion, error) {
	return b.promotionRepo.FindAll(ctx)
}

func (b *PromotionBusiness) FindById(ctx context.Context, id string) (*entity.Promotion, error) {
	return b.promotionRepo.FindById(ctx, id)
}

func (b *PromotionBusiness) FindByCode(ctx context.Context, code string) (*entity.Promotion, error) {
	return b.promotionRepo.FindByCode(ctx, code)
}
