package config

import (
<<<<<<< HEAD
	"backend/models"
=======
	"fmt"
>>>>>>> 54962e854190a18c51ae3ee07cac22c5c8e940dc

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

<<<<<<< HEAD
func InitDB() {
	var err error
	dsn := "root:kozato321@tcp(127.0.0.1:3306)/miniproject?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Admin{})
	DB.AutoMigrate(&models.Redeem{})
}
=======
func IntialDatabase() {
	InitDB()
	// InitMigrate()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {

	config := Config{
		DB_Username: "root",
		DB_Password: "gromizk123",
		DB_Port:     "3306",
		DB_Host:     "db",
		DB_Name:     "c_loyal",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

// func InitMigrate() {
// 	DB.AutoMigrate(&models.User{})
// 	DB.AutoMigrate(&models.Product{})
// 	DB.AutoMigrate(&models.Admin{})
// 	DB.AutoMigrate(&models.Redeem{})
// }
>>>>>>> 54962e854190a18c51ae3ee07cac22c5c8e940dc
