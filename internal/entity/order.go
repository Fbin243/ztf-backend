package entity

type Order struct {
	*BaseEntity
	MerchantId string   `json:"merchant_id"`
	Merchant   Merchant `json:"merchant"`
	UserId     *string  `json:"user_id"`
	User       *User    `json:"user"`
	PayAmount  int64    `json:"pay_amount"`
	Info       string   `json:"info"`
}
