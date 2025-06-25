package composer

import (
	"sync"

	biz "ztf-backend/order/internal/business"
	"ztf-backend/order/internal/repo"

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
	userRepo := repo.NewUserRepo()
	merchantRepo := repo.NewMerchantRepo()
	orderRepo := repo.NewOrderRepo()
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
		c.PromotionConn.Close()
	}
}
