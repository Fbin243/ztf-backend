package tidb

import (
	"context"
	"time"
	"ztf-backend/services/promotion/internal/entity"
	"ztf-backend/services/promotion/internal/errors"

	"gorm.io/gorm"
)

func (r *UserPromotionRepo) UpsertOne(
	ctx context.Context,
	userPromotion *entity.UserPromotion,
) (int64, int64, error) {
	err := r.WithContext(ctx).Save(userPromotion).Error
	if err != nil {
		return 0, 0, err
	}
	return userPromotion.UserId, userPromotion.PromotionId, nil
}

func (r *UserPromotionRepo) DeleteOne(
	ctx context.Context,
	userId int64,
	promotionId int64,
) (int64, int64, error) {
	err := r.WithContext(ctx).
		Delete(&entity.UserPromotion{}, "user_id = ? AND promotion_id = ?", userId, promotionId).
		Error
	if err != nil {
		return 0, 0, err
	}
	return userId, promotionId, nil
}

func (r *UserPromotionRepo) MarkAsUsed(ctx context.Context, req *entity.MarkAsUsedReq) error {
	result := r.WithContext(ctx).Model(&entity.UserPromotion{}).
		Where("user_id = ? AND promotion_id = ? AND used_count < ?", req.UserId, req.PromotionId, entity.MaxUsedCount).
		Updates(map[string]any{
			"used_count":   gorm.Expr("used_count + ?", 1),
			"last_used_at": time.Now(),
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.ErrorNoRowsAffected
	}
	return nil
}
