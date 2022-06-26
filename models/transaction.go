package models

import (
	"time"

	"gorm.io/gorm"
)

//history transaction point (- userid, productid)
type Transaction struct {
	Id        int            `gorm:"column:id" json:"id"`
	UserId    int            `json:"-"`
	ProductId int            `json:"-"`
	User      UserPayload    `json:"user" gorm:"foreignKey:UserId;references:Id"`
	Product   ProductPayload `json:"product" gorm:"foreignKey:ProductId;references:Id"`
	Total     int            `gorm:"column:total" json:"total"`
	Point     int            `gorm:"column:point" json:"point"`
	Date      time.Time
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type InsertTransaction struct {
	Id      int `gorm:"column:id" json:"id"`
	User    int `json:"user" gorm:"column:user_id"`
	Product int `json:"product" gorm:"column:product_id"`
	Total   int `gorm:"column:total" json:"total"`
	Point   int `gorm:"column:point" json:"point"`
}

func (Transaction) TableName() string {
	return "transaction"
}

func (ProductPayload) TableName() string {
	return "product"
}
