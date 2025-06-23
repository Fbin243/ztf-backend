package biz

import (
	"ztf-backend/internal/entity"

	"github.com/google/uuid"
)

func (b *CouponBusiness) FindAll() ([]entity.Coupon, error) {
	return b.couponRepo.FindAll()
}

func (b *CouponBusiness) FindById(id uuid.UUID) (*entity.Coupon, error) {
	return b.couponRepo.FindById(id)
}
