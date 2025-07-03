package tidb

import (
	"log"
	"ztf-backend/services/promotion/internal/entity"

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
