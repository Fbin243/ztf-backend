package entity

type CreateOrderInput struct {
	MerchantId string `json:"merchant_id" validate:"required"`
	Amount     int64  `json:"amount"  validate:"required"`
	Info       string `json:"info"`
}

type UpdateOrderInput struct {
	Amount *int64  `json:"amount"`
	Info   *string `json:"info"`
}

type PayOrderInput struct {
	UserId          string  `json:"user_id"      validate:"required,uuid"`
	PromotionId     *string `json:"promotion_id" validate:"omitempty,uuid"`
	Amount          int64   `json:"amount"       validate:"required"`
	PromotionAmount int64   `json:"promotion_amount" validate:"required"`
	PayAmount       int64   `json:"pay_amount"       validate:"required"`
}
