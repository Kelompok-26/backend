package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id            int            `gorm:"column:id" json:"id"`
	Name          string         `gorm:"column:name" json:"name"`
	Email         string         `gorm:"column:email" json:"email"`
	PhoneNumber   string         `gorm:"column:phone_number" json:"phone_number"`
	Password      string         `gorm:"column:password" json:"password"`
	DateofBirth   time.Time      `gorm:"column:date_of_birth" json:"date_of_birth"`
	Gender        string         `gorm:"column:gender" json:"gender"`
	Point         int            `gorm:"column:point" json:"point"`
	AccountNumber string         `gorm:"column:account_number" json:"account_number"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type UserPayload struct {
	Id   int    `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}
