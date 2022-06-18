package models

import (
	"time"

	"gorm.io/gorm"
)

//history transaction point (- userid, productid)
type transaction struct {
	TransactionId int            `gorm:"column:transactionid" json:"transactionid"`
	Total         int            `gorm:"column:total" json:"total"`
	Point         int            `gorm:"column:point" json:"point"`
	Date          time.Time      
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
