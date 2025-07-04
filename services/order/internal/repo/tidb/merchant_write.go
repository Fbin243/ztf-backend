package tidb

import (
	"context"
	"ztf-backend/services/order/internal/entity"

	"github.com/google/uuid"
)

func (r *MerchantRepo) InsertMany(
	ctx context.Context,
	merchants []entity.Merchant,
) ([]string, error) {
	if len(merchants) == 0 {
		return nil, nil
	}

	for i := range merchants {
		if merchants[i].Id == "" {
			merchants[i].Id = uuid.New().String()
		}
	}

	if err := r.DB.WithContext(ctx).Create(&merchants).Error; err != nil {
		return nil, err
	}

	ids := make([]string, len(merchants))
	for i, merchant := range merchants {
		ids[i] = merchant.Id
	}
	return ids, nil
}
