package repo

import (
	"errors"
	"gorm.io/gorm"
	"ztf-backend/order/internal/entity"
	error2 "ztf-backend/shared/errors"
	"ztf-backend/shared/pkg/db/base"
)

type CouponRepo struct {
	*base.BaseRepo[entity.Coupon]
}

func NewCouponRepo() *CouponRepo {
	return &CouponRepo{base.NewBaseRepo[entity.Coupon]()}
}

func (r *CouponRepo) FindByCode(code string) (*entity.Coupon, error) {
	var coupon entity.Coupon
	err := r.DB.First(&coupon, "code = ?", code).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, error2.ErrorNotFound
	}
	if err != nil {
		return nil, err
	}

	return &coupon, nil
}
