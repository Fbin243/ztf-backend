package repo

import (
	"ztf-backend/internal/entity"
	"ztf-backend/internal/utils"
)

func (r *OrderRepo) UpdateUserId(id string, userId string) (string, error) {
	result := r.DB.Model(&entity.Order{}).Where("id = ? AND user_id IS NULL", id).Update("user_id", userId)
	if result.Error != nil {
		return "", result.Error
	}
	if result.RowsAffected == 0 {
		return "", utils.ErrorNoRowsAffected
	}

	return id, nil
}
