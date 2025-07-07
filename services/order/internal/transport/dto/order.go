package dto

import "time"

type Order struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	PayAmount int64     `json:"pay_amount"`
	Info      string    `json:"info"`
	Merchant  Merchant  `json:"merchant"`
	User      *User     `json:"user"`
}
