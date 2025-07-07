package tidb

import (
	"context"
	"errors"
	"ztf-backend/services/order/internal/entity"
	"ztf-backend/services/order/pkg/idgen"

	"gorm.io/gorm"

	errs "ztf-backend/services/order/internal/errors"
)

func (r *OrderRepo) InsertOne(ctx context.Context, order *entity.Order) (int64, error) {
	order.Id = idgen.GenerateID()
	if err := r.DB.WithContext(ctx).Create(order).Error; err != nil {
		return 0, err
	}
	return order.Id, nil
}

func (r *OrderRepo) InsertMany(ctx context.Context, orders []entity.Order) ([]int64, error) {
	if len(orders) == 0 {
		return nil, nil
	}

	for i := range orders {
		orders[i].Id = idgen.GenerateID()
	}

	if err := r.DB.WithContext(ctx).Create(&orders).Error; err != nil {
		return nil, err
	}

	ids := make([]int64, len(orders))
	for i, order := range orders {
		ids[i] = order.Id
	}
	return ids, nil
}

func (r *OrderRepo) UpdateOne(ctx context.Context, order *entity.Order) (int64, error) {
	if err := r.DB.WithContext(ctx).Save(order).Error; err != nil {
		return 0, err
	}
	return order.Id, nil
}

func (r *OrderRepo) DeleteOne(ctx context.Context, id int64) (int64, error) {
	var order entity.Order
	err := r.DB.WithContext(ctx).Delete(&order, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, errs.ErrorNotFound
	}
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *OrderRepo) UpdatePaymentInfo(
	ctx context.Context,
	id int64,
	order *entity.Order,
) (int64, error) {
	result := r.DB.WithContext(ctx).Model(&entity.Order{}).
		Where("id = ? AND user_id IS NULL", id).
		Updates(order)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errs.ErrorNoRowsAffected
	}

	return id, nil
}
