package biz

import (
	"errors"

	"github.com/jinzhu/copier"
	"ztf-backend/internal/entity"
	"ztf-backend/internal/utils"
)

func (b *OrderBusiness) InsertOne(input *entity.CreateOrderInput) (string, error) {
	// Check if the merchant exists
	exists, err := b.merchantRepo.Exists(input.MerchantId)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", utils.ErrorNotFound
	}

	newOrder := &entity.Order{
		BaseEntity: &entity.BaseEntity{},
	}
	err = copier.Copy(newOrder, input)
	if err != nil {
		return "", err
	}

	return b.orderRepo.InsertOne(newOrder)
}

func (b *OrderBusiness) UpdateOne(id string, input *entity.UpdateOrderInput) (string, error) {
	// Get the existing order
	existingOrder, err := b.orderRepo.FindById(id)
	if err != nil {
		return "", err
	}

	// Check if order is already paid
	if existingOrder.UserId != nil {
		return "", errors.New("order is already paid")
	}

	// Check if the user exists
	if input.UserId != nil {
		exists, err := b.userRepo.Exists(*input.UserId)
		if err != nil {
			return "", err
		}
		if !exists {
			return "", utils.ErrorNotFound
		}
		existingOrder.UserId = input.UserId
	}

	err = copier.Copy(existingOrder, input)
	if err != nil {
		return "", err
	}

	return b.orderRepo.UpdateOne(existingOrder)
}

func (b *OrderBusiness) DeleteOne(id string) (string, error) {
	// Check if the order exists
	exists, err := b.orderRepo.Exists(id)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", utils.ErrorNotFound
	}

	return b.orderRepo.DeleteOne(id)
}
