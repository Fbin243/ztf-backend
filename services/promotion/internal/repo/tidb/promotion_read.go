package tidb

import (
	"context"

	"ztf-backend/services/promotion/internal/entity"
)

func (r *PromotionRepo) FindByCode(ctx context.Context, code string) (*entity.Promotion, error) {
	var promotion entity.Promotion
	err := r.WithContext(ctx).Where("code = ?", code).First(&promotion).Error
	if err != nil {
		return nil, err
	}
	return &promotion, nil
}
