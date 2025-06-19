package biz

import "ztf-backend/internal/entity"

func (b *OrderBusiness) InsertOne(order *entity.Order) (uint, error) {
	return b.orderRepo.InsertOne(order)
}

func (b *OrderBusiness) UpdateOne(order *entity.Order) (uint, error) {
	return b.orderRepo.UpdateOne(order)
}

func (b *OrderBusiness) DeleteOne(id uint) (uint, error) {
	return b.orderRepo.DeleteOne(id)
}
