package business

import "ztf-backend/internal/entity"

type ICouponRepo interface {
	findAll() ([]entity.Coupon, error)
	findById(id string) (*entity.Coupon, error)
	insertOne(coupon *entity.Coupon) (string, error)
	updateOne(coupon *entity.Coupon) (string, error)
	deleteOne(id string) (string, error)
}

type CouponBusiness struct {
	couponRepo ICouponRepo
}

func NewCouponBusiness(couponRepo ICouponRepo) *CouponBusiness {
	return &CouponBusiness{
		couponRepo: couponRepo,
	}
}
