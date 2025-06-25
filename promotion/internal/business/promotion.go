package biz

import (
	"ztf-backend/promotion/internal/entity"
)

type IPromotionRepo interface {
	FindAll() ([]entity.Promotion, error)
	FindById(id string) (*entity.Promotion, error)
	InsertOne(promotion *entity.Promotion) (string, error)
	UpdateOne(promotion *entity.Promotion) (string, error)
	DeleteOne(id string) (string, error)
	Exists(id string) (bool, error)
}

type IUserPromotionRepo interface {
	FindByUserId(userId string) ([]entity.UserPromotion, error)
	InsertOne(userPromotion *entity.UserPromotion) (string, error)
	UpdateOne(userPromotion *entity.UserPromotion) (string, error)
	DeleteOne(id string) (string, error)
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
