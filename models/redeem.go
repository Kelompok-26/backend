package models

import (
	"time"

	"gorm.io/gorm"
)

//redeem for point (-user id)
type redeems struct {
	RedeemId  int            `gorm:"column:redeemid" json:"redeemid"`
	Type      string         `gorm:"column:type" json:"type"`
	Name      string         `gorm:"column:name" json:"name"`
	Nominal   int            `gorm:"column:nominal" json:"nominal"`
	Point     int            `gorm:"column:point" json:"point"`
	Date      time.Time      `json:"-"`
	Status    string         `gorm:"column:status" json:"status"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
