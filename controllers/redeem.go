package controllers

import (
	"backend/config"
	"backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GET All User  Redeem Data "GET -> http://127.0.0.1:8080/api/v1/redeems/"
func GETAllusers(c echo.Context) error {
	var redeems []models.Redeem
	if err := config.DB.Find(&redeems).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, redeems)
}

// GET Spesific Data Redeem User "GET -> http://127.0.0.1:8080/api/redeems/:RID"
func GETSpesUser(c echo.Context) error {

	redeemId, err := strconv.Atoi(c.Param("RID"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid redeem id!")
	}
	var redeem models.Redeem

	if err := config.DB.First(&redeem, redeemId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")

	}
	if redeem.ID == 0 {
		return c.String(http.StatusNotFound, "redeem not found")

	}
	return c.JSON(http.StatusOK, redeem)
}

// CREATE New Redeem "POST -> http://127.0.0.1:8080/api/redeems/"
// {
// 		"Username": "",
//     	"Password": "",
//     	"Name": "",
//     	"Email": ""
// }
func CREATEUser(c echo.Context) error {

	redeem := models.Redeem{}
	if err := c.Bind(&redeem).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	fmt.Printf("Before insert: %#v\n", redeem)
	if err := config.DB.Save(&redeem).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, redeem)
}

// EDIT Spesific User Redeem Data "PUT -> http://127.0.0.1:8080/api/redeems/:RID"
func EDITUser(c echo.Context) error {

	redeemId, err := strconv.Atoi(c.Param("RID"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid redeem id!")
	}
	fmt.Println("Isi redeemId", redeemId)
	var redeem models.Redeem
	fmt.Printf("Isi user redeem sebelum select %#v\n", redeem)
	if err := config.DB.First(&redeem, redeemId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if redeem.ID == 0 {
		return c.String(http.StatusNotFound, "redeem not found")
	}

	fmt.Printf("isi user redeem setelah select %#v\n", redeem)
	if err := c.Bind(&redeem).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	fmt.Printf("Isi user redeem setelah bind %#v\n", redeem)
	fmt.Printf("Before Update : %#v\n", redeem)
	if err := config.DB.Save(&redeem).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, redeem)
}

// DELETE user "DELETE -> http://127.0.0.1:1111/api/redeems/:RID"
func DELETEUserRedeem(c echo.Context) error {

	redeemId, err := strconv.Atoi(c.Param("RID"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid redeem id!")
	}

	var redeem models.Redeem
	if err := config.DB.First(&redeem, redeemId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if redeem.ID == 0 {
		return c.String(http.StatusNotFound, "user not found")
	}
	if err := config.DB.Delete(&redeem).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error!")
	}
	return c.JSON(http.StatusOK, redeem)
}
