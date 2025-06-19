package transport

import biz "ztf-backend/internal/business"

type CouponHandler struct {
	couponBusiness *biz.CouponBusiness
}

func NewCouponHandler(CouponBusiness *biz.CouponBusiness) *CouponHandler {
	return &CouponHandler{
		couponBusiness: CouponBusiness,
	}
}
