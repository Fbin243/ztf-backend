package tidb

import (
	"context"
	"ztf-backend/services/order/internal/entity"
	"ztf-backend/services/order/pkg/idgen"
)

func (r *MerchantRepo) InsertMany(
	ctx context.Context,
	merchants []entity.Merchant,
) ([]int64, error) {
	if len(merchants) == 0 {
		return nil, nil
	}

	for i := range merchants {
		if merchants[i].Id == 0 {
			merchants[i].Id = idgen.GenerateID()
		}
	}

	if err := r.DB.WithContext(ctx).Create(&merchants).Error; err != nil {
		return nil, err
	}

	ids := make([]int64, len(merchants))
	for i, merchant := range merchants {
		ids[i] = merchant.Id
	}
	return ids, nil
}
