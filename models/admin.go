package models

type Admin struct {
	Id          int    `json:"id" form:"id"`
	PhoneNumber string `gorm:"column:phone_number" json:"phone_number"`
	Password    string `json:"password" form:"password"`
}
