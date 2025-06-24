package repo

import (
	"ztf-backend/order/internal/entity"
	error2 "ztf-backend/shared/errors"
)

func (r *OrderRepo) UpdateUserId(id string, userId string) (string, error) {
	result := r.DB.Model(&entity.Order{}).
		Where("id = ? AND user_id IS NULL", id).
		Update("user_id", userId)
	if result.Error != nil {
		return "", result.Error
	}
	if result.RowsAffected == 0 {
		return "", error2.ErrorNoRowsAffected
	}

	return id, nil
}
