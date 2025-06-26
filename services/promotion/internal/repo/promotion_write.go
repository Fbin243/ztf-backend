package repo

import (
	"context"

	"ztf-backend/pkg/errors"
	"ztf-backend/services/promotion/internal/entity"

	"gorm.io/gorm"
)

func (r *PromotionRepo) UpdateRemainingCount(ctx context.Context, id string) error {
	result := r.WithContext(ctx).Model(&entity.Promotion{}).
		Where("id = ? AND remaining_count > 0", id).
		Update("remaining_count", gorm.Expr("remaining_count - ?", 1))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.ErrorNoRowsAffected
	}
	return nil
}
