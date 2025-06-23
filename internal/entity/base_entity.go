package entity

import (
	"time"

	"github.com/google/uuid"
)

type IBaseEntity interface {
	GetID() uuid.UUID
}

type BaseEntity struct {
	Id        uuid.UUID `json:"id"         gorm:"type:char(16);primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (e *BaseEntity) GetID() uuid.UUID {
	return e.Id
}
