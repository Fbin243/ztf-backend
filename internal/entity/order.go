package entity

type Order struct {
	*BaseEntity
	PayAmount int64 `json:"pay_amount" gorm:"not null"`
	CouponId  *uint `json:"coupon_id"`
}
