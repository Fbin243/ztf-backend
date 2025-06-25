package entity

type ApplyPromotionReq struct {
	PromotionId     string
	UserId          string
	OrderId         string
	Amount          int64
	PromotionAmount int64
}
