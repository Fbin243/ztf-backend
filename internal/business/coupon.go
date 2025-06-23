package biz

import (
	"ztf-backend/internal/entity"

	"github.com/google/uuid"
)

type ICouponRepo interface {
	FindAll() ([]entity.Coupon, error)
	FindById(id uuid.UUID) (*entity.Coupon, error)
	InsertOne(coupon *entity.Coupon) (uuid.UUID, error)
	UpdateOne(coupon *entity.Coupon) (uuid.UUID, error)
	DeleteOne(id uuid.UUID) (uuid.UUID, error)
	Exists(id uuid.UUID) (bool, error)
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
