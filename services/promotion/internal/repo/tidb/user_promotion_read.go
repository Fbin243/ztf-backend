package tidb

import (
	"context"
	"errors"
	"ztf-backend/services/promotion/internal/entity"

	"gorm.io/gorm"

	errs "ztf-backend/services/promotion/internal/errors"
)

func (r *UserPromotionRepo) Exists(
	ctx context.Context,
	userId int64,
	promotionId int64,
) (bool, error) {
	var count int64
	err := r.WithContext(ctx).
		Model(&entity.UserPromotion{}).
		Where("user_id = ? AND promotion_id = ?", userId, promotionId).
		Count(&count).
		Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserPromotionRepo) FindByUserId(
	ctx context.Context,
	userId int64,
) ([]entity.UserPromotion, error) {
	var userPromotions []entity.UserPromotion
	err := r.WithContext(ctx).Where("user_id = ?", userId).Find(&userPromotions).Error
	if err != nil {
		return nil, err
	}
	return userPromotions, nil
}

func (r *UserPromotionRepo) FindByUserIdAndPromotionId(
	ctx context.Context,
	userId int64,
	promotionId int64,
) (*entity.UserPromotion, error) {
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
