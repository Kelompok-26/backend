package models

import (
	"time"

	"gorm.io/gorm"
)

type Redeem struct {
	Id        int        `gorm:"column:id" json:"id"`
	Type      string     `gorm:"column:type" json:"type"`
	Name      string     `gorm:"column:name" json:"name"`
	Nominal   int        `gorm:"column:nominal" json:"nominal"`
	Point     int        `gorm:"column:point" json:"point"`
	UserId    int        `json:"-"`
	User      UserRedeem `json:"user" gorm:"foreignKey:UserId;references:Id"`
	Date      time.Time
	Status    string         `gorm:"column:status" json:"status"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type InsertRedeem struct {
	Type    string `gorm:"column:type" json:"type"`
	Name    string `gorm:"column:name" json:"name"`
	Nominal int    `gorm:"column:nominal" json:"nominal"`
	Point   int    `gorm:"column:point" json:"point"`
	Status  string `gorm:"column:status" json:"status"`
	User    int    `json:"user" gorm:"column:user_id"`
}

func (Redeem) TableName() string {
	return "redeem"
}

func (UserRedeem) TableName() string {
	return "user"
}
