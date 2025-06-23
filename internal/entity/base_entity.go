package entity

import "time"

type IBaseEntity interface {
	GetID() uint
}

type BaseEntity struct {
	Id        uint      `json:"id"         gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (e *BaseEntity) GetID() uint {
	return e.Id
}
