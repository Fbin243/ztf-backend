package repo

import (
	"log"

	"ztf-backend/promotion/internal/entity"
	"ztf-backend/shared/pkg/db"
	"ztf-backend/shared/pkg/db/base"
)

type PromotionRepo struct {
	*base.BaseRepo[entity.Promotion]
}

func NewPromotionRepo() *PromotionRepo {
	err := db.GetDB().AutoMigrate(&entity.Promotion{})
	if err != nil {
		log.Printf("Error migrating promotion table: %v", err)
	}

	return &PromotionRepo{base.NewBaseRepo[entity.Promotion]()}
}
