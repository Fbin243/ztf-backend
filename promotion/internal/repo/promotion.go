package repo

import (
	"log"

	"ztf-backend/pkg/db/base"
	biz "ztf-backend/promotion/internal/business"
	"ztf-backend/promotion/internal/entity"

	"gorm.io/gorm"
)

type PromotionRepo struct {
	*base.BaseRepo[entity.Promotion]
}

func NewPromotionRepo(db *gorm.DB) *PromotionRepo {
	err := db.AutoMigrate(&entity.Promotion{})
	if err != nil {
		log.Printf("Error migrating promotion table: %v", err)
	}

	return &PromotionRepo{base.NewBaseRepo[entity.Promotion](db)}
}

func (r *PromotionRepo) WithTx(tx *gorm.DB) biz.IPromotionRepo {
	return NewPromotionRepo(tx)
}
