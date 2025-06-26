package repo

import (
	"ztf-backend/pkg/db/base"
	"ztf-backend/services/order/internal/entity"

	"gorm.io/gorm"
)

type UserRepo struct {
	*base.BaseRepo[entity.User]
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{base.NewBaseRepo[entity.User](db)}
}
