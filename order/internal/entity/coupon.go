package entity

import (
	"time"
	"ztf-backend/shared/pkg/db/base"
)

type CouponType string

const (
	CouponTypePercentage CouponType = "percentage"
	CouponTypeFixed      CouponType = "fixed"
)

type UsageMethod string

const (
	UsageMethodSingleUse UsageMethod = "manual"
	UsageMethodMultiUse  UsageMethod = "automatic"
)

type Coupon struct {
	*base.BaseEntity
	Code           string      `json:"code"             gorm:"not null"`
	Name           string      `json:"name"             gorm:"not null"`
	Value          float64     `json:"value"            gorm:"not null"`
	Description    string      `json:"description"`
	CouponType     CouponType  `json:"coupon_type"      gorm:"not null"`
	UsageMethod    UsageMethod `json:"usage_method"     gorm:"not null"`
	ExpirationDate time.Time   `json:"expiration_date"`
	Orders         []Order     `json:"orders,omitempty" gorm:"foreignKey:CouponId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
