package composer

import (
	"log"
	"sync"

	"ztf-backend/pkg/db"
	biz "ztf-backend/services/order/internal/business"
	"ztf-backend/services/order/internal/repo"

	"google.golang.org/grpc"
)

type Composer struct {
	OrderRepo       biz.IOrderRepo
	UserRepo        biz.IUserRepo
	MerchantRepo    biz.IMerchantRepo
	PromotionClient biz.IPromotionClient

	OrderBusiness    *biz.OrderBusiness
	MerchantBusiness *biz.MerchantBusiness
	UserBusiness     *biz.UserBusiness

	PromotionConn *grpc.ClientConn
}

var (
	composer *Composer
	once     sync.Once
)

func GetComposer() *Composer {
	once.Do(func() {
		db := db.GetDB()
		userRepo := repo.NewUserRepo(db)
		merchantRepo := repo.NewMerchantRepo(db)
		orderRepo := repo.NewOrderRepo(db)
		promotionClient, conn := ComposePromotionClient()

		orderBusiness := biz.NewOrderBusiness(orderRepo, userRepo, merchantRepo, promotionClient)
		merchantBusiness := biz.NewMerchantBusiness(merchantRepo)
		userBusiness := biz.NewUserBusiness(userRepo)

		composer = &Composer{
			OrderRepo:       orderRepo,
			UserRepo:        userRepo,
			MerchantRepo:    merchantRepo,
			PromotionClient: promotionClient,

			OrderBusiness:    orderBusiness,
			MerchantBusiness: merchantBusiness,
			UserBusiness:     userBusiness,

			PromotionConn: conn,
		}
	})
	return composer
}

func (c *Composer) Close() {
	if c.PromotionConn != nil {
		err := c.PromotionConn.Close()
		if err != nil {
			log.Printf("Error closing promotion connection: %v", err)
		}
	}
}
