package biz

import (
	"ztf-backend/internal/entity"
	"ztf-backend/internal/utils"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
)

func (b *OrderBusiness) InsertOne(input *entity.CreateOrderInput) (uuid.UUID, error) {
	newOrder := &entity.Order{}

	// Get the coupon if it exists
	if input.CouponCode != nil {
		coupon, err := b.couponRepo.FindByCode(*input.CouponCode)
		if err != nil {
			return uuid.Nil, err
		}

		newOrder.CouponId = lo.ToPtr(coupon.Id)
	}

	err := copier.Copy(newOrder, input)
	if err != nil {
		return uuid.Nil, err
	}

	return b.orderRepo.InsertOne(newOrder)
}

func (b *OrderBusiness) UpdateOne(id uuid.UUID, input *entity.UpdateOrderInput) (uuid.UUID, error) {
	// Get the existing order
	existingOrder, err := b.orderRepo.FindById(id)
	if err != nil {
		return uuid.Nil, err
	}

	// Get the coupon if it exists
	if input.CouponCode != nil {
		coupon, err := b.couponRepo.FindByCode(*input.CouponCode)
		if err != nil {
			return uuid.Nil, err
		}

		existingOrder.CouponId = lo.ToPtr(coupon.Id)
	}

	err = copier.Copy(existingOrder, input)
	if err != nil {
		return uuid.Nil, err
	}

	return b.orderRepo.UpdateOne(existingOrder)
}

func (b *OrderBusiness) DeleteOne(id uuid.UUID) (uuid.UUID, error) {
	// Check if the order exists
	exists, err := b.orderRepo.Exists(id)
	if err != nil {
		return uuid.Nil, err
	}
	if !exists {
		return uuid.Nil, utils.ErrorNotFound
	}

	return b.orderRepo.DeleteOne(id)
}
