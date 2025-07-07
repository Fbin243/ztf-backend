package entity

import (
	"time"
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
	Id             int64         `json:"id"              gorm:"primaryKey"`
	CreatedAt      time.Time     `json:"created_at"      gorm:"autoCreateTime"`
	UpdatedAt      time.Time     `json:"updated_at"      gorm:"autoUpdateTime"`
	Code           string        `json:"code"            gorm:"not null"`
	Name           string        `json:"name"            gorm:"not null"`
	Value          int64         `json:"value"           gorm:"not null"`
	Description    string        `json:"description"`
	PromotionType  PromotionType `json:"promotion_type"  gorm:"not null"`
	UsageMethod    UsageMethod   `json:"usage_method"    gorm:"not null"`
	ExpirationDate time.Time     `json:"expiration_date"`
	CampaignId     int64         `json:"campaign_id"     gorm:"not null"`
	RemainingCount int64         `json:"remaining_count" gorm:"not null"`
	IsForAll       bool          `json:"is_for_all"      gorm:"not null"`
	// Metadata       map[string]any `json:"metadata"`
}

func (p *Promotion) CalculatePromotionAmount(amount int64) int64 {
	if p.PromotionType == PromotionTypePercentage {
		return amount * p.Value / 100
	}
	return p.Value
}
