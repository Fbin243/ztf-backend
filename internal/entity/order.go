package entity

type Order struct {
	*BaseEntity
	PayAmount  int64  `gorm:"not null" json:"pay_amount"`
	CouponId   uint   `json:"coupon_id"`
}
