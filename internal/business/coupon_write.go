package biz

import (
	"ztf-backend/internal/entity"

	"github.com/jinzhu/copier"
)

func (b *CouponBusiness) InsertOne(input *entity.CreateCouponInput) (uint, error) {
	newCoupon := &entity.Coupon{}
	err := copier.Copy(newCoupon, input)
	if err != nil {
		return 0, err
	}

	return b.couponRepo.InsertOne(newCoupon)
}

func (b *CouponBusiness) UpdateOne(id uint, input *entity.UpdateCouponInput) (uint, error) {
	// Get existing coupon by id
	existingCoupon, err := b.couponRepo.FindById(id)
	if err != nil {
		return 0, err
	}

	err = copier.Copy(existingCoupon, input)
	if err != nil {
		return 0, err
	}

	return b.couponRepo.UpdateOne(existingCoupon)
}

func (b *CouponBusiness) DeleteOne(id uint) (uint, error) {
	// Check if the coupon exists
	_, err := b.couponRepo.FindById(id)
	if err != nil {
		return 0, err
	}

	return b.couponRepo.DeleteOne(id)
}
