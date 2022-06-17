package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId        int           `gorm:"column:userid" json:"userid"`
	Username      string        `gorm:"column:username" json:"username"`
	Email         string        `gorm:"column:email" json:"email"`
	PhoneNumber   int           `gorm:"column:phonenumber" json:"phonenumber"`
	Password      string        `gorm:"column:password" json:"password"`
	DateofBirth   time.Time     `json:"-"`
	Point         int           `gorm:"column:point" json:"point"`
	AccountNumber int           `gorm:"column:accountnumber" json:"accountnumber"`
	Transactions  []Transaction `gorm:"foreignKey:UserId"`
	Redeems       []Redeem      `gorm:"foreignKey:UserId"`
}

// func (User) TableName() string {
// 	return "users"
// }

type Transaction struct {
	gorm.Model
	UserId        int
	TransactionId int       `gorm:"column:transactionid" json:"transactionid"`
	Products      []Product `gorm:"foreignKey:TransactionId"`
	Date          time.Time `json:"-"`
}

type Product struct {
	gorm.Model
	TransactionId int
	ProductId     int    `gorm:"column:productid" json:"productid"`
	TypeProduct   string `gorm:"column:typeproduct" json:"typeproduct"`
	ProviderName  string `gorm:"column:providername" json:"providername"`
	ProductName   string `gorm:"column:productname" json:"productname"`
	Nominal       int    `gorm:"column:nominal" json:"nominal"`
	Point         int    `gorm:"column:point" json:"point"`
	Stock         int    `gorm:"column:stock" json:"stock"`
}

type Redeem struct {
	gorm.Model
	UserId int
	// RedeemId int       `gorm:"column:redeemid" json:"redeemid"`
	Type    string    `gorm:"column:type" json:"type"`
	Name    string    `gorm:"column:name" json:"name"`
	Nominal int       `gorm:"column:nominal" json:"nominal"`
	Point   int       `gorm:"column:point" json:"point"`
	Date    time.Time `json:"-"`
	Status  string    `gorm:"column:status" json:"status"`
}
