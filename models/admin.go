package models

type Admin struct {
	ID          int    `json:"id" form:"id"`
	PhoneNumber int    `gorm:"column:phonenumber" json:"phonenumber"`
	Password    string `json:"password" form:"password"`
}
