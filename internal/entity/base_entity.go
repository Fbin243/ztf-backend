package entity

import "time"

type IBaseEntity interface {
	GetID() uint
}

type BaseEntity struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (e *BaseEntity) GetID() uint {
	return e.Id
}
