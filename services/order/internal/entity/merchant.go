package entity

import "time"

type Merchant struct {
	Id        string    `json:"id"         gorm:"type:char(36);primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Username  string    `json:"username"   gorm:"unique"`
	Email     string    `json:"email"      gorm:"unique"`
}
