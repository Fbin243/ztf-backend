package composer

import (
	"log"
	"sync"

	biz "ztf-backend/order/internal/business"
	"ztf-backend/order/internal/repo"
	"ztf-backend/pkg/db"

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
	db := db.GetDB()
	userRepo := repo.NewUserRepo(db)
	merchantRepo := repo.NewMerchantRepo(db)
	orderRepo := repo.NewOrderRepo(db)
	promotionClient, conn := ComposePromotionClient()

	orderBusiness := biz.NewOrderBusiness(orderRepo, userRepo, merchantRepo, promotionClient)
	merchantBusiness := biz.NewMerchantBusiness(merchantRepo)
	userBusiness := biz.NewUserBusiness(userRepo)

	once.Do(func() {
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
