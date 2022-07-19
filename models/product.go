package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id           int            `gorm:"column:id" json:"id"`
	TypeProduct  string         `gorm:"column:type_product" json:"type_product"`
	ProviderName string         `gorm:"column:provider_name" json:"provider_name"`
	ProductName  string         `gorm:"column:product_name" json:"product_name"`
	Nominal      int            `gorm:"column:nominal" json:"nominal"`
	Point        int            `gorm:"column:point" json:"point"`
	Stock        int            `gorm:"column:stock" json:"stock"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
