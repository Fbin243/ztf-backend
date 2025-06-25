package repo

import (
	"log"

	"ztf-backend/pkg/db"
	"ztf-backend/promotion/internal/entity"

	"gorm.io/gorm"
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
