package biz

import (
	"context"

	"ztf-backend/order/internal/entity"
)

type IUserRepo interface {
	Exists(id string) (bool, error)
	FindByIds(ids []string) ([]entity.User, error)
}

type IMerchantRepo interface {
	Exists(id string) (bool, error)
	FindByIds(ids []string) ([]entity.Merchant, error)
}

type IOrderRepo interface {
	Exists(id string) (bool, error)
	FindAll() ([]entity.Order, error)
	FindById(id string) (*entity.Order, error)
	FindByIds(ids []string) ([]entity.Order, error)
	InsertOne(order *entity.Order) (string, error)
	UpdateOne(order *entity.Order) (string, error)
	DeleteOne(id string) (string, error)
	FindByIdWithMerchantAndUser(id string) (*entity.Order, error)
	UpdatePaymentInfo(id string, order *entity.Order) (string, error)
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
