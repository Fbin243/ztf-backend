package repo

import (
	"ztf-backend/order/internal/entity"
	"ztf-backend/shared/pkg/db/base"
)

type UserRepo struct {
	*base.BaseRepo[entity.User]
}

func NewUserRepo() *UserRepo {
	return &UserRepo{base.NewBaseRepo[entity.User]()}
}
