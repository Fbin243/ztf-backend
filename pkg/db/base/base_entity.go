package base

import (
	"time"
)

type IBaseEntity interface {
	GetId() string
	SetId(id string)
}

type BaseEntity struct {
	Id        string    `json:"id"         gorm:"type:char(36);primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (e *BaseEntity) GetId() string {
	return e.Id
}

func (e *BaseEntity) SetId(id string) {
	e.Id = id
}
