package biz

import (
	"context"

	"ztf-backend/order/internal/entity"
)

type UserBusiness struct {
	userRepo IUserRepo
}

func NewUserBusiness(userRepo IUserRepo) *UserBusiness {
	return &UserBusiness{userRepo: userRepo}
}

func (b *UserBusiness) FindByIds(ctx context.Context, ids []string) ([]entity.User, error) {
	return b.userRepo.FindByIds(ctx, ids)
}

func (b *UserBusiness) ValidateUser(ctx context.Context, userId string) (bool, error) {
	// Check if the user exists
	exists, err := b.userRepo.Exists(ctx, userId)
	if err != nil {
		return false, err
	}
	return exists, nil
}
