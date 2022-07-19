package models

import (
	"time"

	"gorm.io/gorm"
)

//history transaction point (- userid, productid)
type Transaction struct {
	Id        int            `gorm:"column:id" json:"id"`
	UserId    int            `json:"user_id"`
	ProductId int            `json:"product_id"`
	User      User           `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Product   Product        `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Number    string         `json:"number"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}


func (Transaction) TableName() string {
	return "transaction"
}
