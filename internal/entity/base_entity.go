package entity

import (
	"time"
)

type IBaseEntity interface {
	GetID() string
	SetID(id string)
}

type BaseEntity struct {
	Id        string    `json:"id"         gorm:"type:char(36);primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (e *BaseEntity) GetID() string {
	return e.Id
}

func (e *BaseEntity) SetID(id string) {
	e.Id = id
}
