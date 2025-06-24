package entity

import (
	"time"

	"ztf-backend/shared/pkg/db/base"
)

type PromotionType string

const (
	PromotionTypePercentage PromotionType = "percentage"
	PromotionTypeFixed      PromotionType = "fixed"
)

type UsageMethod string

const (
	UsageMethodSingleUse UsageMethod = "manual"
	UsageMethodMultiUse  UsageMethod = "automatic"
)

type Promotion struct {
	*base.BaseEntity
	Code           string        `json:"code"             gorm:"not null"`
	Name           string        `json:"name"             gorm:"not null"`
	Value          float64       `json:"value"            gorm:"not null"`
	Description    string        `json:"description"`
	PromotionType  PromotionType `json:"promotion_type"   gorm:"not null"`
	UsageMethod    UsageMethod   `json:"usage_method"     gorm:"not null"`
	ExpirationDate time.Time     `json:"expiration_date"`
}
