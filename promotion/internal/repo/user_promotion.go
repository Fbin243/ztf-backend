package repo

import (
	"log"

	"gorm.io/gorm"
	"ztf-backend/pkg/db"
	"ztf-backend/promotion/internal/entity"
)

type UserPromotionRepo struct {
	*gorm.DB
}

func NewUserPromotionRepo() *UserPromotionRepo {
	err := db.GetDB().AutoMigrate(&entity.UserPromotion{})
	if err != nil {
		log.Printf("Error migrating user promotion table: %v", err)
	}

	return &UserPromotionRepo{db.GetDB()}
}
