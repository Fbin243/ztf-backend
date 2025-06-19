package biz

import "ztf-backend/internal/entity"

func (b *CouponBusiness) FindAll() ([]entity.Coupon, error) {
	return b.couponRepo.FindAll()
}

func (b *CouponBusiness) FindById(id string) (*entity.Coupon, error) {
	return b.couponRepo.FindById(id)
}
