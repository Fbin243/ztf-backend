package business

import "ztf-backend/internal/entity"

func (b *OrderBusiness) InsertOne(order *entity.Order) (string, error) {
	return b.orderRepo.insertOne(order)
}

func (b *OrderBusiness) UpdateOne(order *entity.Order) (string, error) {
	return b.orderRepo.updateOne(order)
}

func (b *OrderBusiness) DeleteOne(id string) (string, error) {
	return b.orderRepo.deleteOne(id)
}
