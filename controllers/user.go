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
	var users []models.Users
	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all users", users))
}

// get user by id
func Getusercontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.Users{}
	if err := config.DB.First(&user, id).Error; err != nil {
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
func Createusercontrollers(c echo.Context) error {
	user := models.Users{}
	c.Bind(&user)

	if err := config.DB.Debug().Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new user", user))
}

// delete user by id
func Deleteusercontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.Users{}
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "user not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("user deleted successfully", user))
}

// update user by id
func Updateusercontrollers(c echo.Context) error {
	id := c.Param("id")
	user := models.Users{}

	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
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
	if err := config.DB.Where("id = ?", id).Debug().Save(&user).Debug().Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success update user", user))
}
