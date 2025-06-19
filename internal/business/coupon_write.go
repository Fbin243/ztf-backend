package business

import "ztf-backend/internal/entity"

func (b *CouponBusiness) InsertOne(coupon *entity.Coupon) (string, error) {
	return b.couponRepo.insertOne(coupon)
}

func (b *CouponBusiness) UpdateOne(coupon *entity.Coupon) (string, error) {
	return b.couponRepo.updateOne(coupon)
}

func (b *CouponBusiness) DeleteOne(id string) (string, error) {
	return b.couponRepo.deleteOne(id)
}
