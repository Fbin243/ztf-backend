package entity

type MarkAsUsedReq struct {
	UserId      string `json:"user_id"`
	PromotionId string `json:"promotion_id"`
}
