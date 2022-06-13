package models

type Admin struct {
	ID          int    `json:"id" form:"id"`
	PhoneNumber int    `gorm:"column:nomor_hp" json:"nomor_hp"`
	Password    string `json:"password" form:"password"`
}
