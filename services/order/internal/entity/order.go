package entity

import "ztf-backend/pkg/db/base"

type Order struct {
	*base.BaseEntity
	MerchantId      string   `json:"merchant_id"`
	Merchant        Merchant `json:"merchant"`
	UserId          *string  `json:"user_id"`
	User            *User    `json:"user"`
	Info            string   `json:"info"`
	PromotionId     *string  `json:"promotion_id"`
	Amount          int64    `json:"amount"`
	PromotionAmount int64    `json:"promotion_amount"`
	PayAmount       int64    `json:"pay_amount"`
}
