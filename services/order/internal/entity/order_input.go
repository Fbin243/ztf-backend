package entity

type CreateOrderInput struct {
	MerchantId int64  `json:"merchant_id" validate:"required"`
	Amount     int64  `json:"amount"      validate:"required"`
	Info       string `json:"info"`
}

type UpdateOrderInput struct {
	Amount *int64  `json:"amount"`
	Info   *string `json:"info"`
}

type PayOrderInput struct {
	PromotionId     *int64 `json:"promotion_id"     validate:"omitempty,uuid"`
	Amount          int64  `json:"amount"           validate:"min=0"`
	PromotionAmount int64  `json:"promotion_amount" validate:"min=0"`
	PayAmount       int64  `json:"pay_amount"       validate:"min=0"`
}
