package tidb

import (
	"context"
	"errors"
	"ztf-backend/services/order/internal/entity"

	"gorm.io/gorm"

	errs "ztf-backend/services/order/internal/errors"
)

func (r *OrderRepo) FindAll(ctx context.Context) ([]entity.Order, error) {
	var orders []entity.Order
	if err := r.DB.WithContext(ctx).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepo) FindById(ctx context.Context, id int64) (*entity.Order, error) {
	var order entity.Order
	err := r.DB.WithContext(ctx).First(&order, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errs.ErrorNotFound
	}
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepo) FindByIds(ctx context.Context, ids []int64) ([]entity.Order, error) {
	var orders []entity.Order
	if err := r.DB.WithContext(ctx).Where("id IN (?)", ids).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepo) Exists(ctx context.Context, id int64) (bool, error) {
	var count int64
	var order entity.Order
	err := r.DB.WithContext(ctx).Model(&order).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *OrderRepo) FindByIdWithMerchantAndUser(
	ctx context.Context,
	id int64,
) (*entity.Order, error) {
	var order entity.Order
	if err := r.DB.WithContext(ctx).Preload("Merchant").Preload("User").Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
