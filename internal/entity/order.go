package entity

import (
	"github.com/google/uuid"
)

type Order struct {
	*BaseEntity
	PayAmount int64      `json:"pay_amount" gorm:"not null"`
	CouponId  *uuid.UUID `json:"coupon_id" gorm:"type:char(36);"`
}
