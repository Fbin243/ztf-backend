package biz

import (
	"ztf-backend/internal/entity"
	"ztf-backend/internal/utils"

	"github.com/jinzhu/copier"
	"github.com/samber/lo"
)

func (b *OrderBusiness) InsertOne(input *entity.CreateOrderInput) (uint, error) {
	newOrder := &entity.Order{}

	// Get the coupon if it exists
	if input.CouponCode != nil {
		coupon, err := b.couponRepo.FindByCode(*input.CouponCode)
		if err != nil {
			return 0, err
		}

		newOrder.CouponId = lo.ToPtr(coupon.Id)
	}

	err := copier.Copy(newOrder, input)
	if err != nil {
		return 0, err
	}

	return b.orderRepo.InsertOne(newOrder)
}

func (b *OrderBusiness) UpdateOne(id uint, input *entity.UpdateOrderInput) (uint, error) {
	// Get the existing order
	existingOrder, err := b.orderRepo.FindById(id)
	if err != nil {
		return 0, err
	}

	// Get the coupon if it exists
	if input.CouponCode != nil {
		coupon, err := b.couponRepo.FindByCode(*input.CouponCode)
		if err != nil {
			return 0, err
		}

		existingOrder.CouponId = lo.ToPtr(coupon.Id)
	}

	err = copier.Copy(existingOrder, input)
	if err != nil {
		return 0, err
	}

	return b.orderRepo.UpdateOne(existingOrder)
}

func (b *OrderBusiness) DeleteOne(id uint) (uint, error) {
	// Check if the order exists
	exists, err := b.orderRepo.Exists(id)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, utils.ErrorNotFound
	}

	return b.orderRepo.DeleteOne(id)
}
