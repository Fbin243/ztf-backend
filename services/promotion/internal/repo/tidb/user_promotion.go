package tidb

import (
	"log"
	"ztf-backend/services/promotion/internal/entity"

	biz "ztf-backend/services/promotion/internal/business"

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
