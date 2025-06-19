package transport

import (
	"fmt"
	"strconv"

	biz "ztf-backend/internal/business"
	"ztf-backend/internal/entity"

	"github.com/gin-gonic/gin"
)

type CouponHandler struct {
	couponBusiness *biz.CouponBusiness
}

func NewCouponHandler(CouponBusiness *biz.CouponBusiness) *CouponHandler {
	return &CouponHandler{
		couponBusiness: CouponBusiness,
	}
}

func (hdl *CouponHandler) GetAllCoupons(ctx *gin.Context) {
	coupons, err := hdl.couponBusiness.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve coupons"})
		return
	}
	ctx.JSON(200, coupons)
}

func (hdl *CouponHandler) GetCouponById(ctx *gin.Context) {
	id := ctx.Param("id")
	coupon, err := hdl.couponBusiness.FindById(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Coupon not found"})
		return
	}
	ctx.JSON(200, coupon)
}

func (hdl *CouponHandler) CreateCoupon(ctx *gin.Context) {
	var coupon entity.Coupon
	if err := ctx.ShouldBindJSON(&coupon); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	id, err := hdl.couponBusiness.InsertOne(&coupon)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create coupon"})
		return
	}
	ctx.JSON(201, gin.H{"id": id})
}

func (hdl *CouponHandler) UpdateCoupon(ctx *gin.Context) {
	var coupon entity.Coupon
	if err := ctx.ShouldBindJSON(&coupon); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	id, err := hdl.couponBusiness.UpdateOne(&coupon)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update coupon"})
		return
	}
	ctx.JSON(200, gin.H{"id": id})
}

func (hdl *CouponHandler) DeleteCoupen(ctx *gin.Context) {
	id := ctx.Param("id")

	val, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	deletedId, err := hdl.couponBusiness.DeleteOne(uint(val))
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete order"})
		return
	}
	ctx.JSON(200, gin.H{"deleted_id": deletedId})
}
