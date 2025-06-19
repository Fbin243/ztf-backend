package repo

import (
	"ztf-backend/internal/db"
	"ztf-backend/internal/entity"
)

type CouponRepo struct {
	*db.BaseRepo[entity.Coupon]
}

func NewCouponRepo() *CouponRepo {
	return &CouponRepo{}
}
