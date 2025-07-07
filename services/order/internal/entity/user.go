package entity

import "time"

type User struct {
	Id        int64     `json:"id"         gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Username  string    `json:"username"   gorm:"unique"`
	Email     string    `json:"email"      gorm:"unique"`
}
