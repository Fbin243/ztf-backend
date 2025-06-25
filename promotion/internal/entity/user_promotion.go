package entity

import "time"

type UserPromotion struct {
	UserId      string     `json:"user_id"      gorm:"primaryKey;not null"`
	PromotionId string     `json:"promotion_id" gorm:"primaryKey;not null"`
	UsedCount   int64      `json:"used_count"   gorm:"not null"`
	LastUsedAt  *time.Time `json:"last_used_at"`
}
