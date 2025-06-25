package repo

import (
	"context"
	"errors"

	errs "ztf-backend/pkg/errors"
	"ztf-backend/promotion/internal/entity"

	"gorm.io/gorm"
)

func (r *UserPromotionRepo) Exists(ctx context.Context, userId string, promotionId string) (bool, error) {
	var count int64
	err := r.WithContext(ctx).Model(&entity.UserPromotion{}).Where("user_id = ? AND promotion_id = ?", userId, promotionId).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserPromotionRepo) FindByUserId(ctx context.Context, userId string) ([]entity.UserPromotion, error) {
	var userPromotions []entity.UserPromotion
	err := r.WithContext(ctx).Where("user_id = ?", userId).Find(&userPromotions).Error
	if err != nil {
		return nil, err
	}
	return userPromotions, nil
}

func (r *UserPromotionRepo) FindByUserIdAndPromotionId(ctx context.Context, userId string, promotionId string) (*entity.UserPromotion, error) {
	userPromotion := &entity.UserPromotion{}
	err := r.WithContext(ctx).
		Where("user_id = ? AND promotion_id = ?", userId, promotionId).
		First(userPromotion).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errs.ErrorNotFound
	}
	if err != nil {
		return nil, err
	}

	return userPromotion, nil
}
