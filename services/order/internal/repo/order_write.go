package repo

import (
	"context"

	errs "ztf-backend/pkg/errors"
	"ztf-backend/services/order/internal/entity"
)

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
