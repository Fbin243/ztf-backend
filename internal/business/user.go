package biz

import "ztf-backend/internal/entity"

type UserBusiness struct {
	userRepo IUserRepo
}

func NewUserBusiness(userRepo IUserRepo) *UserBusiness {
	return &UserBusiness{userRepo: userRepo}
}

func (b *UserBusiness) FindByIds(ids []string) ([]entity.User, error) {
	return b.userRepo.FindByIds(ids)
}
