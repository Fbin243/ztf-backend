package biz

import (
	"ztf-backend/promotion/internal/entity"
	errs "ztf-backend/shared/errors"

	"github.com/jinzhu/copier"
)

func (b *PromotionBusiness) InsertOne(input *entity.CreatePromotionInput) (string, error) {
	newPromotion := &entity.Promotion{}
	err := copier.Copy(newPromotion, input)
	if err != nil {
		return "", err
	}

	return b.promotionRepo.InsertOne(newPromotion)
}

func (b *PromotionBusiness) UpdateOne(id string, input *entity.UpdatePromotionInput) (string, error) {
	// Check if the promotion exists
	existingPromotion, err := b.promotionRepo.FindById(id)
	if err != nil {
		return "", err
	}

	err = copier.Copy(existingPromotion, input)
	if err != nil {
		return "", err
	}

	return b.promotionRepo.UpdateOne(existingPromotion)
}

func (b *PromotionBusiness) DeleteOne(id string) (string, error) {
	// Check if the promotion exists
	exists, err := b.promotionRepo.Exists(id)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", errs.ErrorNotFound
	}

	return b.promotionRepo.DeleteOne(id)
}
