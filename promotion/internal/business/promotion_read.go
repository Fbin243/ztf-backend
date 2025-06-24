package biz

import (
	"ztf-backend/promotion/internal/entity"
)

func (b *PromotionBusiness) FindAll() ([]entity.Promotion, error) {
	return b.promotionRepo.FindAll()
}

func (b *PromotionBusiness) FindById(id string) (*entity.Promotion, error) {
	return b.promotionRepo.FindById(id)
}
