package biz

import (
	"context"

	"ztf-backend/promotion/internal/entity"
)

type IPromotionRepo interface {
	FindAll(ctx context.Context) ([]entity.Promotion, error)
	FindById(ctx context.Context, id string) (*entity.Promotion, error)
	InsertOne(ctx context.Context, promotion *entity.Promotion) (string, error)
	UpdateOne(ctx context.Context, promotion *entity.Promotion) (string, error)
	DeleteOne(ctx context.Context, id string) (string, error)
	Exists(ctx context.Context, id string) (bool, error)
}

type IUserPromotionRepo interface {
	FindByUserId(ctx context.Context, userId string) ([]entity.UserPromotion, error)
	InsertOne(ctx context.Context, userPromotion *entity.UserPromotion) (string, error)
	UpdateOne(ctx context.Context, userPromotion *entity.UserPromotion) (string, error)
	DeleteOne(ctx context.Context, id string) (string, error)
}

type PromotionBusiness struct {
	promotionRepo     IPromotionRepo
	userPromotionRepo IUserPromotionRepo
}

func NewPromotionBusiness(promotionRepo IPromotionRepo) *PromotionBusiness {
	return &PromotionBusiness{
		promotionRepo: promotionRepo,
	}
}
