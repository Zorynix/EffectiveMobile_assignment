package models

type Car struct {
	RegNum  string `json:"regNum" gorm:"primaryKey"`
	Mark    string `json:"mark"`
	Model   string `json:"model"`
	Year    int    `json:"year,omitempty"`
	OwnerID uint   `json:"-"`
	Owner   People `json:"owner" gorm:"foreignKey:OwnerID"`
}
