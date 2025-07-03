package tidb

import (
	"log"
	"ztf-backend/services/promotion/internal/entity"

	"gorm.io/gorm"
)

type UserPromotionRepo struct {
	*gorm.DB
}

func NewUserPromotionRepo(db *gorm.DB) *UserPromotionRepo {
	err := db.AutoMigrate(&entity.UserPromotion{})
	if err != nil {
		log.Printf("Error migrating user promotion table: %v", err)
	}

	return &UserPromotionRepo{db}
}
