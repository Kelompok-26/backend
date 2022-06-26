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
	Stock        int            `gorm:"column:stock" json:"stock"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type ProductPayload struct {
	Id           int    `gorm:"column:id" json:"id"`
	TypeProduct  string `gorm:"column:type_product" json:"type_product"`
	ProviderName string `gorm:"column:provider_name" json:"provider_name"`
	ProductName  string `gorm:"column:product_name" json:"product_name"`
	Nominal      int    `gorm:"column:nominal" json:"nominal"`
}
