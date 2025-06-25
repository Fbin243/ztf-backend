package biz

import (
	"context"

	"ztf-backend/pkg/db/base"
	"ztf-backend/promotion/internal/entity"

	"gorm.io/gorm"
)

type IPromotionRepo interface {
	base.IBaseRepo[entity.Promotion]
	WithTx(tx *gorm.DB) IPromotionRepo
	UpdateRemainingCount(ctx context.Context, id string) error
}

type IUserPromotionRepo interface {
	WithTx(tx *gorm.DB) IUserPromotionRepo
	Exists(ctx context.Context, userId string, promotionId string) (bool, error)
	FindByUserIdAndPromotionId(ctx context.Context, userId string, promotionId string) (*entity.UserPromotion, error)
	FindByUserId(ctx context.Context, userId string) ([]entity.UserPromotion, error)
	UpsertOne(ctx context.Context, userPromotion *entity.UserPromotion) (string, string, error)
	DeleteOne(ctx context.Context, userId string, promotionId string) (string, string, error)
	MarkAsUsed(ctx context.Context, req *entity.MarkAsUsedReq) error
}

type PromotionBusiness struct {
	promotionRepo     IPromotionRepo
	userPromotionRepo IUserPromotionRepo
}

func NewPromotionBusiness(promotionRepo IPromotionRepo, userPromotionRepo IUserPromotionRepo) *PromotionBusiness {
	return &PromotionBusiness{
		promotionRepo:     promotionRepo,
		userPromotionRepo: userPromotionRepo,
	}
}
