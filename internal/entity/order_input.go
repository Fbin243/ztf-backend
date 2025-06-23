package entity

type CreateOrderInput struct {
	PayAmount  int64   `json:"pay_amount"  validate:"required"`
	CouponCode *string `json:"coupon_code"`
}
type UpdateOrderInput struct {
	PayAmount  *int64  `json:"pay_amount"`
	CouponCode *string `json:"coupon_code"`
}
