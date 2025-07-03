package tidb

import (
	"context"
	"errors"
	"ztf-backend/services/promotion/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"

	errs "ztf-backend/services/promotion/internal/errors"
)

func (r *PromotionRepo) UpdateRemainingCount(ctx context.Context, id string) error {
	result := r.WithContext(ctx).Model(&entity.Promotion{}).
		Where("id = ? AND remaining_count > 0", id).
		Update("remaining_count", gorm.Expr("remaining_count - ?", 1))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errs.ErrorNoRowsAffected
	}
	return nil
}

func (r *PromotionRepo) InsertOne(ctx context.Context, promo *entity.Promotion) (string, error) {
	promo.Id = uuid.New().String()
	if err := r.DB.WithContext(ctx).Create(promo).Error; err != nil {
		return "", err
	}
	return promo.Id, nil
}

func (r *PromotionRepo) InsertMany(
	ctx context.Context,
	promos []entity.Promotion,
) ([]string, error) {
	if len(promos) == 0 {
		return nil, nil
	}

	for i := range promos {
		promos[i].Id = uuid.New().String()
	}

	if err := r.DB.WithContext(ctx).Create(&promos).Error; err != nil {
		return nil, err
	}

	ids := make([]string, len(promos))
	for i, promo := range promos {
		ids[i] = promo.Id
	}
	return ids, nil
}

func (r *PromotionRepo) UpdateOne(ctx context.Context, promo *entity.Promotion) (string, error) {
	if err := r.DB.WithContext(ctx).Save(promo).Error; err != nil {
		return "", err
	}
	return promo.Id, nil
}

func (r *PromotionRepo) DeleteOne(ctx context.Context, id string) (string, error) {
	var promo entity.Promotion
	err := r.DB.WithContext(ctx).Delete(&promo, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errs.ErrorNotFound
	}
	if err != nil {
		return "", err
	}
	return id, nil
}
