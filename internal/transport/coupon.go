package transport

import (
	"net/http"

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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve coupons"})
		return
	}

	ctx.JSON(http.StatusOK, coupons)
}

func (hdl *CouponHandler) GetCouponById(ctx *gin.Context) {
	stringId := ctx.Param("id")
	uuidId, err := utils.ConvertStringToUUID(stringId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	coupon, err := hdl.couponBusiness.FindById(uuidId)
	if err == utils.ErrorNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Coupon not found"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve coupon"})
		return
	}

	ctx.JSON(http.StatusOK, coupon)
}

func (hdl *CouponHandler) CreateCoupon(ctx *gin.Context) {
	var coupon entity.CreateCouponInput
	if err := ctx.ShouldBindJSON(&coupon); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate the coupon data
	err := GetValidator().Struct(coupon)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := hdl.couponBusiness.InsertOne(&coupon)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create coupon"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

func (hdl *CouponHandler) UpdateCoupon(ctx *gin.Context) {
	stringID := ctx.Param("id")
	uuidID, err := utils.ConvertStringToUUID(stringID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var coupon entity.UpdateCouponInput
	if err := ctx.ShouldBindJSON(&coupon); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	id, err := hdl.couponBusiness.UpdateOne(uuidID, &coupon)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update coupon"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (hdl *CouponHandler) DeleteCoupon(ctx *gin.Context) {
	stringId := ctx.Param("id")
	uuidId, err := utils.ConvertStringToUUID(stringId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	id, err := hdl.couponBusiness.DeleteOne(uuidId)
	if err == utils.ErrorNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Coupon not found"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete coupon"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}
