package transport

import (
	biz "ztf-backend/internal/business"
	"ztf-backend/internal/entity"
	"ztf-backend/internal/utils"

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
	stringId := ctx.Param("id")
	uintId, err := utils.ConvertStringToUInt(stringId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	coupon, err := hdl.couponBusiness.FindById(uintId)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Coupon not found"})
		return
	}

	ctx.JSON(200, coupon)
}

func (hdl *CouponHandler) CreateCoupon(ctx *gin.Context) {
	var coupon entity.CreateCouponInput
	if err := ctx.ShouldBindJSON(&coupon); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Validate the coupon data
	err := GetValidator().Struct(coupon)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Validation failed", "details": err.Error()})
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
	stringID := ctx.Param("id")
	uintID, err := utils.ConvertStringToUInt(stringID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	var coupon entity.UpdateCouponInput
	if err := ctx.ShouldBindJSON(&coupon); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	id, err := hdl.couponBusiness.UpdateOne(uintID, &coupon)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update coupon"})
		return
	}
	ctx.JSON(200, gin.H{"id": id})
}

func (hdl *CouponHandler) DeleteCoupen(ctx *gin.Context) {
	stringId := ctx.Param("id")
	uintId, err := utils.ConvertStringToUInt(stringId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	id, err := hdl.couponBusiness.DeleteOne(uintId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete order"})
		return
	}
	ctx.JSON(200, gin.H{"deleted_id": id})
}
