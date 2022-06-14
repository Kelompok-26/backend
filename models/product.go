package models

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	ProductId    int            `gorm:"column:productid" json:"productid"`
	TypeProduct  string         `gorm:"column:typeproduct" json:"typeproduct"`
	ProviderName string         `gorm:"column:providername" json:"providername"`
	ProductName  string         `gorm:"column:productname" json:"productname"`
	Nominal      int            `gorm:"column:nominal" json:"nominal"`
	Point        int            `gorm:"column:point" json:"point"`
	Stock        int            `gorm:"column:stock" json:"stock"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
