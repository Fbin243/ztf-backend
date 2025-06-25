package rpc

import (
	"context"

	"ztf-backend/order/internal/entity"
	"ztf-backend/proto/pb/promotion"
)

type PromotionClient struct {
	promotion.PromotionServiceClient
}

func NewPromotionClient(promotionClient promotion.PromotionServiceClient) *PromotionClient {
	return &PromotionClient{
		promotionClient,
	}
}

func (c *PromotionClient) VerifyPromotion(ctx context.Context, req *entity.VerifyPromotionReq) (bool, error) {
	response, err := c.PromotionServiceClient.VerifyPromotion(ctx, &promotion.VerifyPromotionRequest{
		PromotionId:     req.PromotionId,
		UserId:          req.UserId,
		OrderId:         req.OrderId,
		Amount:          req.Amount,
		PromotionAmount: req.PromotionAmount,
	})
	if err != nil {
		return false, err
	}

	return response.Verified, nil
}
