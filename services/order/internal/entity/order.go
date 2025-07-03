package entity

import "time"

type Order struct {
	Id              string    `json:"id"               gorm:"type:char(36);primaryKey"`
	CreatedAt       time.Time `json:"created_at"       gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at"       gorm:"autoUpdateTime"`
	MerchantId      string    `json:"merchant_id"`
	Merchant        Merchant  `json:"merchant"`
	UserId          *string   `json:"user_id"`
	User            *User     `json:"user"`
	Info            string    `json:"info"`
	PromotionId     *string   `json:"promotion_id"`
	Amount          int64     `json:"amount"`
	PromotionAmount int64     `json:"promotion_amount"`
	PayAmount       int64     `json:"pay_amount"`
}
