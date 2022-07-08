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
func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	// if err := config.DB.Where("PhoneNumber = ? AND password = ?", user.PhoneNumber, user.Password).First(&user).Error; err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }
	if err := config.DB.Table("user").Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// phoneNumber, _ := strconv.Atoi(user.PhoneNumber)
	token, err := middleware.CreateToken(user.Id, "user")
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
	if err := config.DB.Table("user").First(&user, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "user not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Table("user").Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	user.Password, _ = CreateHash(user.Password)

	return c.JSON(http.StatusOK, helper.BuildResponse("user deleted successfully", user))
}

// EDIT Spesific User Data "PUT -> http://127.0.0.1:8080/users/:id"

func UpdateUserControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid user id!")
	}
	fmt.Println("Isi id", id)
	var user models.User
	fmt.Printf("Isi user sebelum select %#v\n", user)
	if err := config.DB.First(&user, id).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if user.Id == 0 {
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
	user.Password, _ = CreateHash(user.Password)

	return c.JSON(http.StatusOK, user)
}

// func UpdateUserControllers(c echo.Context) error {
// 	id := c.Param("id")
// 	fmt.Println(id)
// 	user := models.User{}

// 	if err := config.DB.Table("user").Where("id = ?", id).Find(&user).Error; err != nil {
// 		if err.Error() == "record not found" {
// 			return c.JSON(http.StatusNotFound, map[string]interface{}{
// 				"message": "user not found",
// 			})
// 		}

// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}

// 	newuser := models.User{}
// 	c.Bind(&newuser)

// 	user.Id, _ = strconv.Atoi(id)
// 	fmt.Println(user.Id)
// 	user.Name = newuser.Name
// 	user.Email = newuser.Email
// 	user.Password, _ = CreateHash(newuser.Password)
// 	user.Gender = newuser.Gender
// 	user.DateofBirth = newuser.DateofBirth
// 	user.PhoneNumber = newuser.PhoneNumber
// 	user.AccountNumber = newuser.AccountNumber
// 	user.Point = newuser.Point
// 	if err := config.DB.Table("user").Where("id = ?", id).Debug().Updates(&user).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, helper.BuildResponse("success update user", user))
// }

func CreateHash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}
