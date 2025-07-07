package entity

type ApplyPromotionReq struct {
	PromotionId     int64 `json:"promotion_id"`
	UserId          int64 `json:"user_id"`
	OrderId         int64 `json:"order_id"`
	Amount          int64 `json:"amount"`
	PromotionAmount int64 `json:"promotion_amount"`
}
