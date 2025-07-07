package tidb

import (
	"context"
	"errors"
	"ztf-backend/services/promotion/internal/entity"

	"gorm.io/gorm"

	errs "ztf-backend/services/promotion/internal/errors"
)

func (r *PromotionRepo) FindAll(ctx context.Context) ([]entity.Promotion, error) {
	var promos []entity.Promotion
	if err := r.DB.WithContext(ctx).Find(&promos).Error; err != nil {
		return nil, err
	}
	return promos, nil
}

func (r *PromotionRepo) FindById(ctx context.Context, id int64) (*entity.Promotion, error) {
	var promo entity.Promotion
	err := r.DB.WithContext(ctx).First(&promo, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errs.ErrorNotFound
	}
	if err != nil {
		return nil, err
	}
	return &promo, nil
}

func (r *PromotionRepo) FindByIds(ctx context.Context, ids []int64) ([]entity.Promotion, error) {
	var promos []entity.Promotion
	if err := r.DB.WithContext(ctx).Where("id IN (?)", ids).Find(&promos).Error; err != nil {
		return nil, err
	}
	return promos, nil
}

func (r *PromotionRepo) Exists(ctx context.Context, id int64) (bool, error) {
	var count int64
	var promo entity.Promotion
	err := r.DB.WithContext(ctx).Model(&promo).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *PromotionRepo) FindByCode(ctx context.Context, code string) (*entity.Promotion, error) {
	var promotion entity.Promotion
	err := r.WithContext(ctx).Where("code = ?", code).First(&promotion).Error
	if err != nil {
		return nil, err
	}
	return &promotion, nil
}
