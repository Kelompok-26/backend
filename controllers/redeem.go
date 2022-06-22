package controllers

import (
	"backend/config"
<<<<<<< HEAD
	"backend/helper"
	"backend/models"
=======
	"backend/models"
	"fmt"
>>>>>>> 897dead70a1e5602ffbabddda4d34e91a82bcb19
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

<<<<<<< HEAD
func GetAllRedeemControllers(c echo.Context) error {
	var redeem []models.Redeem
	if err := config.DB.Model(&redeem).Debug().Preload("User").Find(&redeem).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all redeem", redeem))
}

func GetRedeemByUserIDControllers(c echo.Context) error {
	user_id, _ := strconv.Atoi(c.Param("user_id"))
	redeem := []models.Redeem{}
	if err := config.DB.Model(&redeem).Debug().Where("user_id", user_id).Preload("User").Find(&redeem).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "redeem not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get redeem", redeem))
}

func CreateRedeemsControllers(c echo.Context) error {
	redeem := models.InsertRedeem{}
	c.Bind(&redeem)

	if err := config.DB.Table("redeem").Create(&redeem).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new redeem", redeem))
}

// delete Redeem by id
func DeleteRedeemControllers(c echo.Context) error {
	id := c.Param("id")
	redeem := models.InsertRedeem{}

	if err := config.DB.Table("redeem").Where("id", id).Find(&redeem).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "redeem not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Table("redeem").Where("id", id).Delete(&redeem).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("redeem deleted successfully", redeem))
}

// update by id
func UpdateRedeemControllers(c echo.Context) error {
	id := c.Param("id")
	redeem := models.InsertRedeem{}

	if err := config.DB.Table("redeem").Where("id", id).Find(&redeem).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "redeem not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	newredeem := models.InsertRedeem{}
	c.Bind(&newredeem)

	redeem.Name = newredeem.Name
	redeem.Nominal = newredeem.Nominal
	redeem.Point = newredeem.Point
	redeem.Status = newredeem.Status
	redeem.Type = newredeem.Type
	redeem.User = newredeem.User
	if err := config.DB.Table("redeem").Where("id", id).Save(&redeem).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success update redeem", redeem))
=======
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
>>>>>>> 897dead70a1e5602ffbabddda4d34e91a82bcb19
}
