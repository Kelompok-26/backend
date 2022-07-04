package controllers

import (
	"backend/config"
	"backend/helper"
	"backend/middleware"
	"backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// LOGIN User "POST -> http://127.0.0.1:8080/login"
func LoginUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	// if err := config.DB.Where("PhoneNumber = ? AND password = ?", user.PhoneNumber, user.Password).First(&user).Error; err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }
	if err := config.DB.Table("user").Where("email = ? OR phone_number = ? AND password = ?", user.Email, user.PhoneNumber, user.Password).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// phoneNumber, _ := strconv.Atoi(user.PhoneNumber)
	token, err := middleware.CreateToken(user.Id, user.Email, "user")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil login",
		"User":    token,
	})

}

// User Regist "POST -> http://127.0.0.1:8080/users
// {
// 		"Name": "",
// 		"Email": "",
//		"PhoneNumber": "",
// 		"Password": "",
//		"DateoBirth": "",
// 		"AccountNumber": ""
// }
func CreateUserControllers(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	user.Password, _ = CreateHash(user.Password)
	if err := config.DB.Table("user").Debug().Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new user", user))
}

//GET All User Data "GET -> http://127.0.0.1:8080/users"
func GetAllusercontrollers(c echo.Context) error {
	var users []models.User
	if err := config.DB.Table("user").Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all users", users))
}

// GET Spesific User Data "GET -> http://127.0.0.1:8080/users/:uid"
func GetUserControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("uid"))
	user := models.User{}
	if err := config.DB.Table("user").First(&user, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "user not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get user", user))
}

// Delete User Data "DELETE -> http://127.0.0.1:8080/users/:uid
func DeleteUserControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("uid"))
	user := models.User{}
	if err := config.DB.Table("user").Where("id = ?", id).First(&user).Error; err != nil {
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

// EDIT Spesific User Data "PUT -> http://127.0.0.1:8080/users/:uid"
func UpdateUserControllers(c echo.Context) error {
	id := c.Param("uid")
	user := models.User{}

	if err := config.DB.Table("user").First(&user, "uid = ?", id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "user not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	newuser := models.User{}
	c.Bind(&newuser)
	fmt.Println("user", user)
	user.Name = newuser.Name
	user.Email = newuser.Email
	user.Password = newuser.Password
	user.DateofBirth = newuser.DateofBirth
	user.PhoneNumber = newuser.PhoneNumber
	user.AccountNumber = newuser.AccountNumber
	user.Point = newuser.Point
	if err := config.DB.Table("user").Where("uid = ?", id).Debug().Save(&user).Debug().Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fmt.Printf("Isi user setelah bind %#v\n", user)
	fmt.Printf("Before Update : %#v\n", user)
	if err := config.DB.Save(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error3") //?
	}

	return c.JSON(http.StatusOK, user)
}

func CreateHash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}
