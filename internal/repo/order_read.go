package repo

import "ztf-backend/internal/entity"

func (r *OrderRepo) FindByIdWithMerchantAndUser(id string) (*entity.Order, error) {
	var order entity.Order
	if err := r.DB.Preload("Merchant").Preload("User").Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
