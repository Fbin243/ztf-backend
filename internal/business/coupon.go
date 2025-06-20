package biz

import "ztf-backend/internal/entity"

type ICouponRepo interface {
	FindAll() ([]entity.Coupon, error)
	FindById(id uint) (*entity.Coupon, error)
	InsertOne(coupon *entity.Coupon) (uint, error)
	UpdateOne(coupon *entity.Coupon) (uint, error)
	DeleteOne(id uint) (uint, error)
}

type CouponBusiness struct {
	couponRepo ICouponRepo
}

func NewCouponBusiness(couponRepo ICouponRepo) *CouponBusiness {
	return &CouponBusiness{
		couponRepo: couponRepo,
	}
}
