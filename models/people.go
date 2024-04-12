package models

type People struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}
