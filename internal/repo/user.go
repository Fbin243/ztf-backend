package repo

import (
	"ztf-backend/internal/db"
	"ztf-backend/internal/entity"
)

type UserRepo struct {
	*db.BaseRepo[entity.User]
}

func NewUserRepo() *UserRepo {
	return &UserRepo{db.NewBaseRepo[entity.User]()}
}
