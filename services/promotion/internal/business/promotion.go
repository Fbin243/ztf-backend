package biz

import (
	"context"
	"ztf-backend/services/promotion/internal/entity"
)

type IPromotionRepo interface {
	FindAll(ctx context.Context) ([]entity.Promotion, error)
	FindById(ctx context.Context, id int64) (*entity.Promotion, error)
	FindByIds(ctx context.Context, ids []int64) ([]entity.Promotion, error)
	InsertOne(ctx context.Context, entity *entity.Promotion) (int64, error)
	InsertMany(ctx context.Context, entities []entity.Promotion) ([]int64, error)
	UpdateOne(ctx context.Context, entity *entity.Promotion) (int64, error)
	DeleteOne(ctx context.Context, id int64) (int64, error)
	Exists(ctx context.Context, id int64) (bool, error)
	UpdateRemainingCount(ctx context.Context, id int64) error
	FindByCode(ctx context.Context, code string) (*entity.Promotion, error)
}

type IUserPromotionRepo interface {
	Exists(ctx context.Context, userId int64, promotionId int64) (bool, error)
	FindByUserIdAndPromotionId(
		ctx context.Context,
		userId int64,
		promotionId int64,
	) (*entity.UserPromotion, error)
	FindByUserId(ctx context.Context, userId int64) ([]entity.UserPromotion, error)
	UpsertOne(ctx context.Context, userPromotion *entity.UserPromotion) (int64, int64, error)
	DeleteOne(ctx context.Context, userId int64, promotionId int64) (int64, int64, error)
	MarkAsUsed(ctx context.Context, req *entity.MarkAsUsedReq) error
}

type IOrderClient interface {
	ValidateUser(ctx context.Context, req *entity.ValidateUserReq) (bool, error)
}

type Tx interface {
	PromotionRepo() IPromotionRepo
	UserPromotionRepo() IUserPromotionRepo
}

type ITxRunner interface {
	Transaction(ctx context.Context, fn func(tx Tx) error) error
}

type PromotionBusiness struct {
	txRunner          ITxRunner
	promotionRepo     IPromotionRepo
	userPromotionRepo IUserPromotionRepo
	orderClient       IOrderClient
}

func NewPromotionBusiness(
	txRunner ITxRunner,
	promotionRepo IPromotionRepo,
	userPromotionRepo IUserPromotionRepo,
	orderClient IOrderClient,
) *PromotionBusiness {
	return &PromotionBusiness{
		txRunner:          txRunner,
		promotionRepo:     promotionRepo,
		userPromotionRepo: userPromotionRepo,
		orderClient:       orderClient,
	}
}
