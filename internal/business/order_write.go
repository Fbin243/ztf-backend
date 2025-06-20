package biz

import (
	"ztf-backend/internal/entity"
	"ztf-backend/internal/utils"

	"github.com/jinzhu/copier"
)

func (b *OrderBusiness) InsertOne(input *entity.CreateOrderInput) (uint, error) {
	newOrder := &entity.Order{}
	err := copier.Copy(newOrder, input)
	if err != nil {
		return 0, err
	}

	return b.orderRepo.InsertOne(newOrder)
}

func (b *OrderBusiness) UpdateOne(id uint, input *entity.UpdateOrderInput) (uint, error) {
	existingOrder, err := b.orderRepo.FindById(id)
	if err != nil {
		return 0, err
	}
	if existingOrder == nil {
		return 0, utils.ErrorNotFound
	}

	err = copier.Copy(existingOrder, input)
	if err != nil {
		return 0, err
	}

	return b.orderRepo.UpdateOne(existingOrder)
}

func (b *OrderBusiness) DeleteOne(id uint) (uint, error) {
	return b.orderRepo.DeleteOne(id)
}
