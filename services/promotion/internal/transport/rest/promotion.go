package rest

import (
	"errors"
	"net/http"
	"ztf-backend/services/promotion/internal/auth"
	"ztf-backend/services/promotion/internal/entity"
	"ztf-backend/services/promotion/internal/transport/validation"

	"github.com/gin-gonic/gin"

	biz "ztf-backend/services/promotion/internal/business"

	errs "ztf-backend/services/promotion/internal/errors"
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
	promotions, err := hdl.promotionBusiness.GetPromotionList(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, promotions)
}

func (hdl *PromotionHandler) GetPromotionById(ctx *gin.Context) {
	stringId := ctx.Param("id")
	promotion, err := hdl.promotionBusiness.GetPromotion(ctx, stringId)
	if errors.Is(err, errs.ErrorNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, promotion)
}

func (hdl *PromotionHandler) GetPromotionByCode(ctx *gin.Context) {
	code := ctx.Query("code")
	promotion, err := hdl.promotionBusiness.GetPromotionByCode(ctx, code)
	if errors.Is(err, errs.ErrorNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, promotion)
}

func (hdl *PromotionHandler) CreatePromotion(ctx *gin.Context) {
	var promotion entity.CreatePromotionInput
	if err := ctx.ShouldBindJSON(&promotion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the promotion data
	err := validation.GetValidator().Struct(promotion)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := hdl.promotionBusiness.CreatePromotion(ctx, &promotion)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if promotion.PromotionType == entity.PromotionTypePercentage &&
		promotion.Value > 100 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Value must be less than or equal to 100%"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

func (hdl *PromotionHandler) UpdatePromotion(ctx *gin.Context) {
	stringID := ctx.Param("id")
	var promotion entity.UpdatePromotionInput
	if err := ctx.ShouldBindJSON(&promotion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validation.GetValidator().Struct(promotion)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := hdl.promotionBusiness.UpdatePromotion(ctx, stringID, &promotion)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (hdl *PromotionHandler) DeletePromotion(ctx *gin.Context) {
	stringId := ctx.Param("id")
	id, err := hdl.promotionBusiness.DeletePromotion(ctx, stringId)
	if errors.Is(err, errs.ErrorNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (hdl *PromotionHandler) VerifyPromotion(ctx *gin.Context) {
	req := &entity.ApplyPromotionReq{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.UserId = ctx.GetHeader("X-User-Id")
	valid, err := hdl.promotionBusiness.VerifyPromotion(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"valid": valid})
}

func (hdl *PromotionHandler) CollectPromotion(ctx *gin.Context) {
	reqCtx := auth.SetAuthKey(ctx, ctx.GetHeader("X-User-Id"))
	promotionId := ctx.Param("id")

	collected, err := hdl.promotionBusiness.CollectPromotion(reqCtx, promotionId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": collected})
}
