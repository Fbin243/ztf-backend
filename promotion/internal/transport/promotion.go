package transport

import (
	"net/http"

	biz "ztf-backend/promotion/internal/business"
	"ztf-backend/promotion/internal/entity"
	errs "ztf-backend/shared/errors"
	"ztf-backend/shared/validation"

	"github.com/gin-gonic/gin"
)

type PromotionHandler struct {
	promotionBusiness *biz.PromotionBusiness
}

func NewPromotionHandler(PromotionBusiness *biz.PromotionBusiness) *PromotionHandler {
	return &PromotionHandler{
		promotionBusiness: PromotionBusiness,
	}
}

func (hdl *PromotionHandler) GetAllPromotions(ctx *gin.Context) {
	promotions, err := hdl.promotionBusiness.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve promotions"})
		return
	}

	ctx.JSON(http.StatusOK, promotions)
}

func (hdl *PromotionHandler) GetPromotionById(ctx *gin.Context) {
	stringId := ctx.Param("id")
	promotion, err := hdl.promotionBusiness.FindById(stringId)
	if err == errs.ErrorNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Promotion not found"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve promotion"})
		return
	}
	ctx.JSON(http.StatusOK, promotion)
}

func (hdl *PromotionHandler) CreatePromotion(ctx *gin.Context) {
	var promotion entity.CreatePromotionInput
	if err := ctx.ShouldBindJSON(&promotion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate the promotion data
	err := validation.GetValidator().Struct(promotion)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := hdl.promotionBusiness.InsertOne(&promotion)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create promotion"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

func (hdl *PromotionHandler) UpdatePromotion(ctx *gin.Context) {
	stringID := ctx.Param("id")
	var promotion entity.UpdatePromotionInput
	if err := ctx.ShouldBindJSON(&promotion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	id, err := hdl.promotionBusiness.UpdateOne(stringID, &promotion)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update promotion"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (hdl *PromotionHandler) DeletePromotion(ctx *gin.Context) {
	stringId := ctx.Param("id")
	id, err := hdl.promotionBusiness.DeleteOne(stringId)
	if err == errs.ErrorNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Promotion not found"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete promotion"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}
