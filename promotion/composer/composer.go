package composer

import (
	"sync"

	biz "ztf-backend/promotion/internal/business"
	"ztf-backend/promotion/internal/repo"
)

type Composer struct {
	PromotionRepo biz.IPromotionRepo

	PromotionBusiness *biz.PromotionBusiness
}

var (
	composer *Composer
	once     sync.Once
)

func GetComposer() *Composer {
	once.Do(func() {
		promotionRepo := repo.NewPromotionRepo()

		promotionBusiness := biz.NewPromotionBusiness(promotionRepo)

		composer = &Composer{
			PromotionRepo:     promotionRepo,
			PromotionBusiness: promotionBusiness,
		}
	})

	return composer
}
