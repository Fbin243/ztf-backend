package tidb

import (
	"context"
	"errors"
	"ztf-backend/services/order/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"

	errs "ztf-backend/services/order/internal/errors"
)

func (r *OrderRepo) InsertOne(ctx context.Context, order *entity.Order) (string, error) {
	order.Id = uuid.New().String()
	if err := r.DB.WithContext(ctx).Create(order).Error; err != nil {
		return "", err
	}
	return order.Id, nil
}

func (r *OrderRepo) InsertMany(ctx context.Context, orders []entity.Order) ([]string, error) {
	if len(orders) == 0 {
		return nil, nil
	}

	for i := range orders {
		orders[i].Id = uuid.New().String()
	}

	if err := r.DB.WithContext(ctx).Create(&orders).Error; err != nil {
		return nil, err
	}

	ids := make([]string, len(orders))
	for i, order := range orders {
		ids[i] = order.Id
	}
	return ids, nil
}

func (r *OrderRepo) UpdateOne(ctx context.Context, order *entity.Order) (string, error) {
	if err := r.DB.WithContext(ctx).Save(order).Error; err != nil {
		return "", err
	}
	return order.Id, nil
}

func (r *OrderRepo) DeleteOne(ctx context.Context, id string) (string, error) {
	var order entity.Order
	err := r.DB.WithContext(ctx).Delete(&order, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errs.ErrorNotFound
	}
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *OrderRepo) UpdatePaymentInfo(
	ctx context.Context,
	id string,
	order *entity.Order,
) (string, error) {
	result := r.DB.WithContext(ctx).Model(&entity.Order{}).
		Where("id = ? AND user_id IS NULL", id).
		Updates(order)
	if result.Error != nil {
		return "", result.Error
	}
	if result.RowsAffected == 0 {
		return "", errs.ErrorNoRowsAffected
	}

	return id, nil
}
