package tidb

import (
	"context"
	"errors"
	"ztf-backend/services/order/internal/entity"

	"gorm.io/gorm"

	errs "ztf-backend/services/order/internal/errors"
)

func (r *MerchantRepo) FindAll(ctx context.Context) ([]entity.Merchant, error) {
	var merchants []entity.Merchant
	if err := r.DB.WithContext(ctx).Find(&merchants).Error; err != nil {
		return nil, err
	}
	return merchants, nil
}

func (r *MerchantRepo) FindById(ctx context.Context, id int64) (*entity.Merchant, error) {
	var merchant entity.Merchant
	err := r.DB.WithContext(ctx).First(&merchant, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errs.ErrorNotFound
	}
	if err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (r *MerchantRepo) FindByIds(ctx context.Context, ids []int64) ([]entity.Merchant, error) {
	var merchants []entity.Merchant
	if err := r.DB.WithContext(ctx).Where("id IN (?)", ids).Find(&merchants).Error; err != nil {
		return nil, err
	}
	return merchants, nil
}

func (r *MerchantRepo) Exists(ctx context.Context, id int64) (bool, error) {
	var count int64
	var merchant entity.Merchant
	err := r.DB.WithContext(ctx).Model(&merchant).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
