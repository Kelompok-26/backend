package models

type Admin struct {
	ID          int    `json:"id" form:"id"`
	PhoneNumber int    `gorm:"column:phone_number" json:"phone_number"`
	Password    string `json:"password" form:"password"`
}
