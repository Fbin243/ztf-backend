package entity

import "time"

type CreatePromotionInput struct {
	Code           string        `json:"code"            validate:"required"`
	Name           string        `json:"name"            validate:"required"`
	Value          float64       `json:"value"           validate:"required"`
	Description    string        `json:"description"`
	PromotionType  PromotionType `json:"promotion_type"  validate:"required,oneof=percentage fixed"`
	UsageMethod    UsageMethod   `json:"usage_method"    validate:"required,oneof=manual automatic"`
	ExpirationDate time.Time     `json:"expiration_date" validate:"required"`
}

type UpdatePromotionInput struct {
	Code           *string        `json:"code"`
	Name           *string        `json:"name"`
	Description    *string        `json:"description"`
	PromotionType  *PromotionType `json:"promotion_type"  validate:"omitempty,oneof=percentage fixed"`
	UsageMethod    *UsageMethod   `json:"usage_method"    validate:"omitempty,oneof=manual automatic"`
	ExpirationDate *time.Time     `json:"expiration_date"`
}
