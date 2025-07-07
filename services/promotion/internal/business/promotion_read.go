package biz

import (
	"context"
	"ztf-backend/services/promotion/internal/entity"
)

func (b *PromotionBusiness) GetPromotionList(ctx context.Context) ([]entity.Promotion, error) {
	return b.promotionRepo.FindAll(ctx)
}

func (b *PromotionBusiness) GetPromotion(
	ctx context.Context,
	id int64,
) (*entity.Promotion, error) {
	return b.promotionRepo.FindById(ctx, id)
}

func (b *PromotionBusiness) GetPromotionByCode(
	ctx context.Context,
	code string,
) (*entity.Promotion, error) {
	return b.promotionRepo.FindByCode(ctx, code)
}
