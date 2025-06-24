package biz

import (
	"errors"

	"ztf-backend/order/internal/entity"
	errs "ztf-backend/shared/errors"
	"ztf-backend/shared/pkg/db/base"

	"github.com/jinzhu/copier"
)

func (b *OrderBusiness) InsertOne(input *entity.CreateOrderInput) (string, error) {
	// Check if the merchant exists
	exists, err := b.merchantRepo.Exists(input.MerchantId)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", errs.ErrorNotFound
	}

	newOrder := &entity.Order{
		BaseEntity: &base.BaseEntity{},
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
		return "", errs.ErrorNotFound
	}

	return b.orderRepo.DeleteOne(id)
}

func (b *OrderBusiness) PayForOrder(id string, input *entity.PayOrderInput) (string, error) {
	// Check if the order exists
	exists, err := b.orderRepo.Exists(id)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", errs.ErrorNotFound
	}

	// Check if the user exists
	exists, err = b.userRepo.Exists(input.UserId)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", errs.ErrorNotFound
	}

	// Pay for the order
	id, err = b.orderRepo.UpdateUserId(id, input.UserId)
	if errors.Is(err, errs.ErrorNoRowsAffected) {
		return "", errors.New("order is already paid")
	}
	if err != nil {
		return "", err
	}

	return id, nil
}
