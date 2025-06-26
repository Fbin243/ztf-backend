package biz

import (
	"context"

	"ztf-backend/pkg/db/base"
	"ztf-backend/services/order/internal/entity"
)

type IUserRepo interface {
	Exists(ctx context.Context, id string) (bool, error)
	FindByIds(ctx context.Context, ids []string) ([]entity.User, error)
}

type IMerchantRepo interface {
	Exists(ctx context.Context, id string) (bool, error)
	FindByIds(ctx context.Context, ids []string) ([]entity.Merchant, error)
}

type IOrderRepo interface {
	base.IBaseRepo[entity.Order]
	FindByIdWithMerchantAndUser(ctx context.Context, id string) (*entity.Order, error)
	UpdatePaymentInfo(ctx context.Context, id string, order *entity.Order) (string, error)
}

type IPromotionClient interface {
	ApplyPromotion(ctx context.Context, req *entity.ApplyPromotionReq) (bool, error)
}

type OrderBusiness struct {
	orderRepo       IOrderRepo
	merchantRepo    IMerchantRepo
	userRepo        IUserRepo
	promotionClient IPromotionClient
}

func NewOrderBusiness(
	orderRepo IOrderRepo,
	userRepo IUserRepo,
	merchantRepo IMerchantRepo,
	promotionClient IPromotionClient,
) *OrderBusiness {
	return &OrderBusiness{
		orderRepo:       orderRepo,
		userRepo:        userRepo,
		merchantRepo:    merchantRepo,
		promotionClient: promotionClient,
	}
}
