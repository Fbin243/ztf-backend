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
	ID             uint        `gorm:"primaryKey" json:"id"`
	CreateAt       time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
	Code           string      `gorm:"not null" json:"code"`
	Name           string      `gorm:"not null" json:"name"`
	Description    string      `json:"description"`
	CouponType     CouponType  `gorm:"not null" json:"coupon_type"`
	UsageMethod    UsageMethod `gorm:"not null" json:"usage_method"`
	ExpirationDate time.Time   `json:"expiration_date"`
}
