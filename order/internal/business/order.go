package biz

import (
	"context"

	"ztf-backend/order/internal/entity"
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
	Exists(ctx context.Context, id string) (bool, error)
	FindAll(ctx context.Context) ([]entity.Order, error)
	FindById(ctx context.Context, id string) (*entity.Order, error)
	FindByIds(ctx context.Context, ids []string) ([]entity.Order, error)
	InsertOne(ctx context.Context, order *entity.Order) (string, error)
	UpdateOne(ctx context.Context, order *entity.Order) (string, error)
	DeleteOne(ctx context.Context, id string) (string, error)
	FindByIdWithMerchantAndUser(ctx context.Context, id string) (*entity.Order, error)
	UpdatePaymentInfo(ctx context.Context, id string, order *entity.Order) (string, error)
}

type IPromotionClient interface {
	VerifyPromotion(ctx context.Context, req *entity.VerifyPromotionReq) (bool, error)
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
