package biz

import (
	"ztf-backend/order/internal/entity"
	errs "ztf-backend/shared/errors"

	"github.com/jinzhu/copier"
)

func (b *CouponBusiness) InsertOne(input *entity.CreateCouponInput) (string, error) {
	newCoupon := &entity.Coupon{}
	err := copier.Copy(newCoupon, input)
	if err != nil {
		return "", err
	}

	return b.couponRepo.InsertOne(newCoupon)
}

func (b *CouponBusiness) UpdateOne(id string, input *entity.UpdateCouponInput) (string, error) {
	// Check if the coupon exists
	existingCoupon, err := b.couponRepo.FindById(id)
	if err != nil {
		return "", err
	}

	err = copier.Copy(existingCoupon, input)
	if err != nil {
		return "", err
	}

	return b.couponRepo.UpdateOne(existingCoupon)
}

func (b *CouponBusiness) DeleteOne(id string) (string, error) {
	// Check if the coupon exists
	exists, err := b.couponRepo.Exists(id)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", errs.ErrorNotFound
	}

	return b.couponRepo.DeleteOne(id)
}
