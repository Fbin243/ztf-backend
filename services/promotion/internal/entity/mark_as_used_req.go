package entity

type MarkAsUsedReq struct {
	UserId      int64 `json:"user_id"`
	PromotionId int64 `json:"promotion_id"`
}
