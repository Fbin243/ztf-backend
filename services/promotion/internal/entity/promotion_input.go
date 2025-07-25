package entity

import "time"

type CreatePromotionInput struct {
	Code           string         `json:"code"            validate:"required,min=5,max=20"`
	Name           string         `json:"name"            validate:"required"`
	Value          int64          `json:"value"           validate:"required,min=0"`
	Description    string         `json:"description"`
	PromotionType  PromotionType  `json:"promotion_type"  validate:"required,oneof=percentage fixed"`
	UsageMethod    UsageMethod    `json:"usage_method"    validate:"required,oneof=manual automatic"`
	ExpirationDate time.Time      `json:"expiration_date" validate:"required"`
	CampaignId     int64          `json:"campaign_id"     validate:"required"`
	Metadata       map[string]any `json:"metadata"`
	IsForAll       bool           `json:"is_for_all"`
	RemainingCount int64          `json:"remaining_count"`
}

type UpdatePromotionInput struct {
	Code           *string         `json:"code"            validate:"omitempty,min=5,max=20"`
	Name           *string         `json:"name"`
	Description    *string         `json:"description"`
	PromotionType  *PromotionType  `json:"promotion_type"  validate:"omitempty,oneof=percentage fixed"`
	UsageMethod    *UsageMethod    `json:"usage_method"    validate:"omitempty,oneof=manual automatic"`
	ExpirationDate *time.Time      `json:"expiration_date"`
	CampaignId     *int64          `json:"campaign_id"     validate:"omitempty"`
	Metadata       *map[string]any `json:"metadata"`
	IsForAll       *bool           `json:"is_for_all"`
	RemainingCount *int64          `json:"remaining_count"`
}
