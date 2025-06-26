package composer

import (
	"sync"

	"google.golang.org/grpc"
	"ztf-backend/pkg/db"
	biz "ztf-backend/services/promotion/internal/business"
	"ztf-backend/services/promotion/internal/repo/rpc"
	"ztf-backend/services/promotion/internal/repo/tidb"
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
		promotionRepo := tidb.NewPromotionRepo(db)
		userPromotionRepo := tidb.NewUserPromotionRepo(db)
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
		_ = c.OrderConn.Close()
	}
}
