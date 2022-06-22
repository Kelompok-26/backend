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

// CREATE New User "POST -> http://127.0.0.1:8080/api/users/"
// {
//     	"Email": "",
// 		"PhoneNumber: ",
// 		"Username: ",
// 		"DateofBirth: ",
// 		"Password: "
// }
func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Debug().Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new user", user))
}

// GET All User Data "GET -> http://127.0.0.1:8080/api/users/"
func GetAllUsers(c echo.Context) error {
	var users []models.User
	// if err:= database.FindAllUsers().Error; err != nil
	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all users", users))
}

// GET Spesific Data User "GET -> http://127.0.0.1:8080/api/users/:UID"
func GetSpesificUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
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

// DELETE user "DELETE -> http://127.0.0.1:8080/api/users/:UID"
func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
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

// EDIT Spesific User Data "PUT -> http://127.0.0.1:1111/api/users/:UID"
func EDITUser(c echo.Context) error {

	userId, err := strconv.Atoi(c.Param("UID"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid user id!")
	}
	fmt.Println("Isi userId", userId)
	var user models.User
	fmt.Printf("Isi user sebelum select %#v\n", user)
	if err := config.DB.First(&user, userId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if user.ID == 0 {
		return c.String(http.StatusNotFound, "user not found")
	}

	fmt.Printf("isi user setelah select %#v\n", user)
	if err := c.Bind(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error2")
	}

	fmt.Printf("Isi user setelah bind %#v\n", user)
	fmt.Printf("Before Update : %#v\n", user)
	if err := config.DB.Save(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error3") //?
	}

	return c.JSON(http.StatusOK, user)
}
