package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	UserId        int            `gorm:"column:userid" json:"userid"`
	Name          string         `gorm:"column:name" json:"name"`
	Email         string         `gorm:"column:email" json:"email"`
	PhoneNumber   int            `gorm:"column:phonenumber" json:"phonenumber"`
	Password      string         `gorm:"column:password" json:"password"`
	DateofBirth   time.Time      `json:"-"`
	Gender        string         `gorm:"column:gender" json:"gender"`
	Point         int            `gorm:"column:point" json:"point"`
	AccountNumber int            `gorm:"column:accountnumber" json:"accountnumber"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
