package tidb

import (
	"context"
	"ztf-backend/services/order/internal/entity"

	"github.com/google/uuid"
)

func (r *UserRepo) InsertMany(ctx context.Context, users []entity.User) ([]string, error) {
	if len(users) == 0 {
		return nil, nil
	}

	for i := range users {
		if users[i].Id == "" {
			users[i].Id = uuid.New().String()
		}
	}

	if err := r.DB.WithContext(ctx).Create(&users).Error; err != nil {
		return nil, err
	}

	ids := make([]string, len(users))
	for i, user := range users {
		ids[i] = user.Id
	}
	return ids, nil
}
