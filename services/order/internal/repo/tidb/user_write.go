package tidb

import (
	"context"
	"ztf-backend/services/order/internal/entity"
	"ztf-backend/services/order/pkg/idgen"
)

func (r *UserRepo) InsertMany(ctx context.Context, users []entity.User) ([]int64, error) {
	if len(users) == 0 {
		return nil, nil
	}

	for i := range users {
		if users[i].Id == 0 {
			users[i].Id = idgen.GenerateID()
		}
	}

	if err := r.DB.WithContext(ctx).Create(&users).Error; err != nil {
		return nil, err
	}

	ids := make([]int64, len(users))
	for i, user := range users {
		ids[i] = user.Id
	}
	return ids, nil
}
