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

type PromotionBusiness struct {
	promotionRepo IPromotionRepo
}

func NewPromotionBusiness(promotionRepo IPromotionRepo) *PromotionBusiness {
	return &PromotionBusiness{
		promotionRepo: promotionRepo,
	}
}
