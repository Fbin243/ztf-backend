package entity

type Order struct {
	*BaseEntity
	PayAmount float64 `gorm:"not null" json:"pay_amount"`
	CouponID  uint    `json:"coupon_id"`
	Coupon    Coupon  `gorm:"foreignKey:CouponID" json:"coupon"`
}
