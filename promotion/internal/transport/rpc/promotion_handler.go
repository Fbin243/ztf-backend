package rpc

import (
	"context"

	biz "ztf-backend/promotion/internal/business"
	"ztf-backend/promotion/internal/entity"
	"ztf-backend/proto/pb/promotion"
)

type PromotionHandler struct {
	promotion.UnimplementedPromotionServiceServer
	promotionBusiness *biz.PromotionBusiness
}

func NewPromotionHandler(promotionBusiness *biz.PromotionBusiness) *PromotionHandler {
	return &PromotionHandler{
		promotionBusiness: promotionBusiness,
	}
}

func (h *PromotionHandler) VerifyPromotion(ctx context.Context, req *promotion.VerifyPromotionRequest) (*promotion.VerifyPromotionResponse, error) {
	res := &promotion.VerifyPromotionResponse{
		Verified: false,
	}

	verified, err := h.promotionBusiness.VerifyPromotion(ctx, &entity.VerifyPromotionReq{
		PromotionId:     req.PromotionId,
		UserId:          req.UserId,
		OrderId:         req.OrderId,
		Amount:          req.Amount,
		PromotionAmount: req.PromotionAmount,
	})

	res.Verified = verified

	return res, err
}
