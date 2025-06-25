package rpc

import (
	"context"
	"fmt"
	"log"

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

func (h *PromotionHandler) ApplyPromotion(
	ctx context.Context,
	req *promotion.ApplyPromotionRequest,
) (*promotion.ApplyPromotionResponse, error) {
	res := &promotion.ApplyPromotionResponse{
		Success: false,
	}

	success, err := h.promotionBusiness.ApplyPromotion(ctx, &entity.ApplyPromotionReq{
		PromotionId:     req.PromotionId,
		UserId:          req.UserId,
		OrderId:         req.OrderId,
		Amount:          req.Amount,
		PromotionAmount: req.PromotionAmount,
	})
	if err != nil {
		return res, fmt.Errorf("failed to apply promotion: %w", err)
	}

	log.Printf("Promotion is applied: %+v", success)
	res.Success = success

	return res, nil
}
