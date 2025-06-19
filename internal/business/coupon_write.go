package biz

import "ztf-backend/internal/entity"

func (b *CouponBusiness) InsertOne(coupon *entity.Coupon) (uint, error) {
	return b.couponRepo.InsertOne(coupon)
}

func (b *CouponBusiness) UpdateOne(coupon *entity.Coupon) (uint, error) {
	return b.couponRepo.UpdateOne(coupon)
}

func (b *CouponBusiness) DeleteOne(id uint) (uint, error) {
	return b.couponRepo.DeleteOne(id)
}
