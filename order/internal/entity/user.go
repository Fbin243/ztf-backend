package entity

import "ztf-backend/pkg/db/base"

type User struct {
	*base.BaseEntity
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email"    gorm:"unique"`
}
