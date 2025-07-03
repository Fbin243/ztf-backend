package composer

import (
	"sync"
	"ztf-backend/services/promotion/internal/repo/rpc"
	"ztf-backend/services/promotion/internal/repo/tidb"

	biz "ztf-backend/services/promotion/internal/business"

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
		db := tidb.GetDB()
		promotionRepo := tidb.NewPromotionRepo(db)
		userPromotionRepo := tidb.NewUserPromotionRepo(db)
		orderClient, conn := ComposeOrderClient()
		txRunner := tidb.NewGormTxRunner(db)

		promotionBusiness := biz.NewPromotionBusiness(
			txRunner,
			promotionRepo,
			userPromotionRepo,
			orderClient,
		)

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
