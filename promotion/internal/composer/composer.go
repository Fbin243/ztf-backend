package composer

import (
	"sync"

	"ztf-backend/pkg/db"
	biz "ztf-backend/promotion/internal/business"
	"ztf-backend/promotion/internal/repo"
)

type Composer struct {
	PromotionRepo     biz.IPromotionRepo
	UserPromotionRepo biz.IUserPromotionRepo

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

		promotionBusiness := biz.NewPromotionBusiness(promotionRepo, userPromotionRepo)

		composer = &Composer{
			PromotionRepo:     promotionRepo,
			UserPromotionRepo: userPromotionRepo,

			PromotionBusiness: promotionBusiness,
		}
	})

	return composer
}
