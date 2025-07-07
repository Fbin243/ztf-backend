package biz

import (
	"context"
	"ztf-backend/services/order/internal/entity"
)

type IUserRepo interface {
	Exists(ctx context.Context, id int64) (bool, error)
	FindByIds(ctx context.Context, ids []int64) ([]entity.User, error)
}

type IMerchantRepo interface {
	Exists(ctx context.Context, id int64) (bool, error)
	FindByIds(ctx context.Context, ids []int64) ([]entity.Merchant, error)
}

type IOrderRepo interface {
	FindAll(ctx context.Context) ([]entity.Order, error)
	FindById(ctx context.Context, id int64) (*entity.Order, error)
	FindByIds(ctx context.Context, ids []int64) ([]entity.Order, error)
	InsertOne(ctx context.Context, entity *entity.Order) (int64, error)
	InsertMany(ctx context.Context, entities []entity.Order) ([]int64, error)
	UpdateOne(ctx context.Context, entity *entity.Order) (int64, error)
	DeleteOne(ctx context.Context, id int64) (int64, error)
	Exists(ctx context.Context, id int64) (bool, error)
	FindByIdWithMerchantAndUser(ctx context.Context, id int64) (*entity.Order, error)
	UpdatePaymentInfo(ctx context.Context, id int64, order *entity.Order) (int64, error)
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
