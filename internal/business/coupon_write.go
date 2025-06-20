package biz

import (
	"log"

	"ztf-backend/internal/entity"
	"ztf-backend/internal/utils"

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
	log.Printf("Updating order with ID: %d", id)
	// Get existing coupon by ID
	existingCoupon, err := b.couponRepo.FindById(id)
	if err != nil {
		return 0, err
	}
	if existingCoupon == nil {
		return 0, utils.ErrorNotFound
	}

	err = copier.Copy(existingCoupon, input)
	if err != nil {
		return 0, err
	}

	log.Printf("Found existing order: %+v", existingCoupon)
	return b.couponRepo.UpdateOne(existingCoupon)
}

func (b *CouponBusiness) DeleteOne(id uint) (uint, error) {
	return b.couponRepo.DeleteOne(id)
}
