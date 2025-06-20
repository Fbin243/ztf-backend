package biz

import "ztf-backend/internal/entity"

func (b *CouponBusiness) FindAll() ([]entity.Coupon, error) {
	return b.couponRepo.FindAll()
}

func (b *CouponBusiness) FindById(id uint) (*entity.Coupon, error) {
	return b.couponRepo.FindById(id)
}
