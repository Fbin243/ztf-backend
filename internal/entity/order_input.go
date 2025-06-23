package entity

type CreateOrderInput struct {
	MerchantId string `json:"merchant_id" validate:"required"`
	PayAmount  int64  `json:"pay_amount"  validate:"required"`
	Info       string `json:"info"`
}

type UpdateOrderInput struct {
	PayAmount *int64  `json:"pay_amount"`
	Info      *string `json:"info"`
	UserId    *string `json:"user_id"`
}
