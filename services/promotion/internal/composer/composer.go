package composer

import (
	"sync"

	"ztf-backend/pkg/db"
	biz "ztf-backend/services/promotion/internal/business"
	"ztf-backend/services/promotion/internal/repo"
	"ztf-backend/services/promotion/internal/repo/rpc"

	"google.golang.org/grpc"
)

type Composer struct {
	PromotionRepo     biz.IPromotionRepo
	UserPromotionRepo biz.IUserPromotionRepo
	OrderClient       *rpc.OrderClient
	OrderConn         *grpc.ClientConn

	PromotionBusiness *biz.PromotionBusiness
}

var (
	composer *Composer
	once     sync.Once
)

func GetComposer() *Composer {
	once.Do(func() {
		db := db.GetDB()
		promotionRepo := repo.NewPromotionRepo(db)
		userPromotionRepo := repo.NewUserPromotionRepo(db)
		orderClient, conn := ComposeOrderClient()

		promotionBusiness := biz.NewPromotionBusiness(promotionRepo, userPromotionRepo, orderClient)

		composer = &Composer{
			PromotionRepo:     promotionRepo,
			UserPromotionRepo: userPromotionRepo,
			OrderClient:       orderClient,
			OrderConn:         conn,

			PromotionBusiness: promotionBusiness,
		}
	})

	return composer
}

func (c *Composer) Close() {
	if c.OrderConn != nil {
		c.OrderConn.Close()
	}
}
