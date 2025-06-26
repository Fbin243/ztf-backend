package tidb

import (
	"gorm.io/gorm"
	"ztf-backend/pkg/db/base"
	"ztf-backend/services/order/internal/entity"
)

type UserRepo struct {
	*base.BaseRepo[entity.User]
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{base.NewBaseRepo[entity.User](db)}
}
