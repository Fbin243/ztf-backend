package entity

import (
	"time"
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
	*BaseEntity
	Code           string      `gorm:"not null" json:"code"`
	Name           string      `gorm:"not null" json:"name"`
	Value          float64     `gorm:"not null" json:"value"`
	Description    string      `json:"description"`
	CouponType     CouponType  `gorm:"not null" json:"coupon_type"`
	UsageMethod    UsageMethod `gorm:"not null" json:"usage_method"`
	ExpirationDate time.Time   `json:"expiration_date"`
	Orders         []Order     `gorm:"foreignKey:CouponId;references:Id" json:"orders,omitempty"`
}
