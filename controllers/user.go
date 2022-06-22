package controllers

import (
	"backend/config"
	"backend/helper"
	"backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all users
func GetAllusercontrollers(c echo.Context) error {
<<<<<<< HEAD
	var users []models.User
	if err := config.DB.Table("user").Find(&users).Error; err != nil {
=======
	var users []models.Users
	if err := config.DB.Find(&users).Error; err != nil {
>>>>>>> 897dead70a1e5602ffbabddda4d34e91a82bcb19
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all users", users))
}

// get user by id
func GetUserControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
<<<<<<< HEAD
	user := models.User{}
	if err := config.DB.Table("user").First(&user, id).Error; err != nil {
=======
	user := models.Users{}
	if err := config.DB.First(&user, id).Error; err != nil {
>>>>>>> 897dead70a1e5602ffbabddda4d34e91a82bcb19
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "user not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get user", user))
}

// create user by id
<<<<<<< HEAD
func CreateUserControllers(c echo.Context) error {
	user := models.User{}
=======
func Createusercontrollers(c echo.Context) error {
	user := models.Users{}
>>>>>>> 897dead70a1e5602ffbabddda4d34e91a82bcb19
	c.Bind(&user)

	if err := config.DB.Table("user").Debug().Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new user", user))
}

// delete user by id
func DeleteUserControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
<<<<<<< HEAD
	user := models.User{}
	if err := config.DB.Table("user").Where("id = ?", id).First(&user).Error; err != nil {
=======
	user := models.Users{}
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
>>>>>>> 897dead70a1e5602ffbabddda4d34e91a82bcb19
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "user not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Table("user").Where("id = ?", id).Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("user deleted successfully", user))
}

// update user by id
func UpdateUserControllers(c echo.Context) error {
	id := c.Param("id")
	user := models.Users{}

	if err := config.DB.Table("user").First(&user, "id = ?", id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "user not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	newuser := models.Users{}
	c.Bind(&newuser)
	fmt.Println("user", user)
	user.Name = newuser.Name
	user.Email = newuser.Email
	user.Password = newuser.Password
	user.Gender = newuser.Gender
	user.DateofBirth = newuser.DateofBirth
	user.PhoneNumber = newuser.PhoneNumber
	user.AccountNumber = newuser.AccountNumber
	user.Point = newuser.Point
	if err := config.DB.Table("user").Where("id = ?", id).Debug().Save(&user).Debug().Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success update user", user))
}
