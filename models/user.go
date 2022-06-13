package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            int    `gorm:"column:id" json:"id"`
	Name          string `gorm:"column:nama" json:"nama"`
	Email         string `gorm:"column:email" json:"email"`
	PhoneNumber   int    `gorm:"column:nomor_hp" json:"nomor_hp"`
	Password      string `gorm:"column:password" json:"password"`
	DateofBirth   time.Time `json:"-"`
	Gender        string `gorm:"column:jenis_kelamin" json:"jenis_kelamin"`
	Point         int    `gorm:"column:point" json:"point"`
	AccountNumber int    `gorm:"column:nomor_akun" json:"nomor_akun"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
