package repo

import (
	"ztf-backend/order/internal/entity"
	"ztf-backend/pkg/db/base"

	"gorm.io/gorm"
)

type UserRepo struct {
	*base.BaseRepo[entity.User]
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{base.NewBaseRepo[entity.User](db)}
}
