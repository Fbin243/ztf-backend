package rpc

import (
	"context"
	"ztf-backend/proto/pb/promotion"
	"ztf-backend/services/order/internal/entity"
)

type PromotionClient struct {
	promotion.PromotionServiceClient
}

func NewPromotionClient(promotionClient promotion.PromotionServiceClient) *PromotionClient {
	return &PromotionClient{
		promotionClient,
	}
}

func (c *PromotionClient) ApplyPromotion(
	ctx context.Context,
	req *entity.ApplyPromotionReq,
) (bool, error) {
	response, err := c.PromotionServiceClient.ApplyPromotion(
		ctx,
		&promotion.ApplyPromotionRequest{
			PromotionId:     req.PromotionId,
			UserId:          req.UserId,
			OrderId:         req.OrderId,
			Amount:          req.Amount,
			PromotionAmount: req.PromotionAmount,
		},
	)
	if err != nil {
		return false, err
	}

	return response.Success, nil
}
