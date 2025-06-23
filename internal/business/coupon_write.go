package biz

import (
	"ztf-backend/internal/entity"
	"ztf-backend/internal/utils"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

func (b *CouponBusiness) InsertOne(input *entity.CreateCouponInput) (uuid.UUID, error) {
	newCoupon := &entity.Coupon{}
	err := copier.Copy(newCoupon, input)
	if err != nil {
		return uuid.Nil, err
	}

	return b.couponRepo.InsertOne(newCoupon)
}

func (b *CouponBusiness) UpdateOne(id uuid.UUID, input *entity.UpdateCouponInput) (uuid.UUID, error) {
	// Check if the coupon exists
	existingCoupon, err := b.couponRepo.FindById(id)
	if err != nil {
		return uuid.Nil, err
	}

	err = copier.Copy(existingCoupon, input)
	if err != nil {
		return uuid.Nil, err
	}

	return b.couponRepo.UpdateOne(existingCoupon)
}

func (b *CouponBusiness) DeleteOne(id uuid.UUID) (uuid.UUID, error) {
	// Check if the coupon exists
	exists, err := b.couponRepo.Exists(id)
	if err != nil {
		return uuid.Nil, err
	}
	if !exists {
		return uuid.Nil, utils.ErrorNotFound
	}

	return b.couponRepo.DeleteOne(id)
}
