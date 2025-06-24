package repo

import (
	"ztf-backend/order/internal/entity"
)

func (r *OrderRepo) FindByIds(ids []string) ([]entity.Order, error) {
	var orders []entity.Order
	if err := r.DB.Where("id IN (?)", ids).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
