package biz

import (
	"ztf-backend/order/internal/entity"
)

type ICouponRepo interface {
	FindAll() ([]entity.Coupon, error)
	FindById(id string) (*entity.Coupon, error)
	InsertOne(coupon *entity.Coupon) (string, error)
	UpdateOne(coupon *entity.Coupon) (string, error)
	DeleteOne(id string) (string, error)
	Exists(id string) (bool, error)
	FindByCode(code string) (*entity.Coupon, error)
}

type CouponBusiness struct {
	couponRepo ICouponRepo
}

func NewCouponBusiness(couponRepo ICouponRepo) *CouponBusiness {
	return &CouponBusiness{
		couponRepo: couponRepo,
	}
}
