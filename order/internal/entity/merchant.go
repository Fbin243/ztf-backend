package entity

import "ztf-backend/shared/pkg/db/base"

type Merchant struct {
	*base.BaseEntity
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email"    gorm:"unique"`
}
