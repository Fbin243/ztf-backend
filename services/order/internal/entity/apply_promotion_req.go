package entity

type ApplyPromotionReq struct {
	PromotionId     int64
	UserId          int64
	OrderId         int64
	Amount          int64
	PromotionAmount int64
}
