package entity

type Merchant struct {
	*BaseEntity
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email"    gorm:"unique"`
}
