package models

type Admin struct {
	Id       int    `json:"id" form:"id"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `json:"password" form:"password"`
}
