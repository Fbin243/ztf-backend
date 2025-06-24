package repo

import (
	"ztf-backend/promotion/internal/entity"
	"ztf-backend/shared/pkg/db/base"
)

type PromotionRepo struct {
	*base.BaseRepo[entity.Promotion]
}

func NewPromotionRepo() *PromotionRepo {
	return &PromotionRepo{base.NewBaseRepo[entity.Promotion]()}
}
