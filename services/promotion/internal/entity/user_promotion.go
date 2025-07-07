package entity

import "time"

type UserPromotion struct {
	UserId      int64      `json:"user_id"      gorm:"primaryKey;not null"`
	PromotionId int64      `json:"promotion_id" gorm:"primaryKey;not null"`
	UsedCount   int64      `json:"used_count"   gorm:"not null"`
	LastUsedAt  *time.Time `json:"last_used_at"`
}
