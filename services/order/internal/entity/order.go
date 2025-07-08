package entity

import "time"

type Order struct {
	Id              int64     `json:"id"               gorm:"primaryKey"`
	CreatedAt       time.Time `json:"created_at"       gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at"       gorm:"autoUpdateTime"`
	MerchantId      int64     `json:"merchant_id"`
	Merchant        Merchant  `json:"merchant"         gorm:"foreignKey:MerchantId;references:Id"`
	UserId          *int64    `json:"user_id"          gorm:"foreignKey:UserId;references:Id"`
	User            *User     `json:"user"`
	Info            string    `json:"info"`
	PromotionId     *int64    `json:"promotion_id"`
	Amount          int64     `json:"amount"`
	PromotionAmount int64     `json:"promotion_amount"`
	PayAmount       int64     `json:"pay_amount"`
}
