package business

import "ztf-backend/internal/entity"

func (b *CouponBusiness) FindAll() ([]entity.Coupon, error) {
	return b.couponRepo.findAll()
}

func (b *CouponBusiness) FindById(id string) (*entity.Coupon, error) {
	return b.couponRepo.findById(id)
}
