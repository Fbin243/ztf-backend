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
