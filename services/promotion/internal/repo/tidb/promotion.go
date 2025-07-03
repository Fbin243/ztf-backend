package tidb

import (
	"log"
	"ztf-backend/services/promotion/internal/entity"

	biz "ztf-backend/services/promotion/internal/business"

	"gorm.io/gorm"
)

type PromotionRepo struct {
	*gorm.DB
}

func NewPromotionRepo(db *gorm.DB) *PromotionRepo {
	err := db.AutoMigrate(&entity.Promotion{})
	if err != nil {
		log.Printf("Error migrating promotion table: %v", err)
	}

	return &PromotionRepo{db}
}

func (r *PromotionRepo) WithTx(tx *gorm.DB) biz.IPromotionRepo {
	return &PromotionRepo{tx}
}
