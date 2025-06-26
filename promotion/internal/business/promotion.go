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
	FindByCode(ctx context.Context, code string) (*entity.Promotion, error)
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

type IOrderClient interface {
	ValidateUser(ctx context.Context, req *entity.ValidateUserReq) (bool, error)
}

type PromotionBusiness struct {
	promotionRepo     IPromotionRepo
	userPromotionRepo IUserPromotionRepo
	orderClient       IOrderClient
}

func NewPromotionBusiness(promotionRepo IPromotionRepo, userPromotionRepo IUserPromotionRepo, orderClient IOrderClient) *PromotionBusiness {
	return &PromotionBusiness{
		promotionRepo:     promotionRepo,
		userPromotionRepo: userPromotionRepo,
		orderClient:       orderClient,
	}
}
