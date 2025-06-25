package repo

import (
	"log"

	biz "ztf-backend/promotion/internal/business"
	"ztf-backend/promotion/internal/entity"

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

func (r *UserPromotionRepo) WithTx(tx *gorm.DB) biz.IUserPromotionRepo {
	return &UserPromotionRepo{tx}
}
