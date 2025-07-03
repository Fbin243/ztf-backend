package tidb

import (
	"gorm.io/gorm"
)

type UserRepo struct {
	*gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}
