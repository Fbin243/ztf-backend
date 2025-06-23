package repo

import (
	"gorm.io/gorm"
	"ztf-backend/internal/db"
	"ztf-backend/internal/entity"
	"ztf-backend/internal/utils"
)

type CouponRepo struct {
	*db.BaseRepo[entity.Coupon]
}

func NewCouponRepo() *CouponRepo {
	return &CouponRepo{db.NewBaseRepo[entity.Coupon]()}
}

func (r *CouponRepo) FindByCode(code string) (*entity.Coupon, error) {
	var coupon entity.Coupon
	err := r.DB.First(&coupon, "code = ?", code).Error
	if err == gorm.ErrRecordNotFound {
		return nil, utils.ErrorNotFound
	}
	if err != nil {
		return nil, err
	}

	return &coupon, nil
}
