package tidb

import (
	"context"

	"ztf-backend/services/order/internal/entity"
)

func (r *OrderRepo) FindByIds(ctx context.Context, ids []string) ([]entity.Order, error) {
	var orders []entity.Order
	if err := r.DB.WithContext(ctx).Where("id IN (?)", ids).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
