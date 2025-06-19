package entity

import "time"

type IBaseEntity interface {
	GetID() uint
}

type BaseEntity struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreateAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (e *BaseEntity) GetID() uint {
	return e.ID
}
