package entity

type ApplyPromotionReq struct {
	PromotionId     string `json:"promotion_id"`
	UserId          string `json:"user_id"`
	OrderId         string `json:"order_id"`
	Amount          int64  `json:"amount"`
	PromotionAmount int64  `json:"promotion_amount"`
}
