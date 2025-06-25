package repo

import (
	"ztf-backend/order/internal/entity"
	errs "ztf-backend/pkg/errors"
)

func (r *OrderRepo) UpdatePaymentInfo(id string, order *entity.Order) (string, error) {
	result := r.DB.Model(&entity.Order{}).
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
