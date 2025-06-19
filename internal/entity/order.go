package entity

import "time"

type Order struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreateAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	PayAmount float64   `gorm:"not null" json:"pay_amount"`
	CouponID  uint      `json:"coupon_id"`
	Coupon    Coupon    `gorm:"foreignKey:CouponID" json:"coupon"`
}
