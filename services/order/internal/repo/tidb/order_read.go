package tidb

import (
	"context"

	"ztf-backend/services/order/internal/entity"
)

func (r *OrderRepo) FindByIdWithMerchantAndUser(
	ctx context.Context,
	id string,
) (*entity.Order, error) {
	var order entity.Order
	if err := r.DB.WithContext(ctx).Preload("Merchant").Preload("User").Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
