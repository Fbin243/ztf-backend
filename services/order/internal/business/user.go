package biz

import (
	"context"
	"ztf-backend/services/order/internal/entity"
)

type UserBusiness struct {
	userRepo IUserRepo
}

func NewUserBusiness(userRepo IUserRepo) *UserBusiness {
	return &UserBusiness{userRepo: userRepo}
}

func (b *UserBusiness) FindByIds(ctx context.Context, ids []int64) ([]entity.User, error) {
	return b.userRepo.FindByIds(ctx, ids)
}

func (b *UserBusiness) ValidateUser(ctx context.Context, userId int64) (bool, error) {
	// Check if the user exists
	exists, err := b.userRepo.Exists(ctx, userId)
	if err != nil {
		return false, err
	}
	return exists, nil
}
