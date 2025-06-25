package repo

import (
	"log"

	"ztf-backend/pkg/db"
	"ztf-backend/pkg/db/base"
	"ztf-backend/promotion/internal/entity"
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
