package controllers

import (
	"backend/config"
	"backend/controllers/request"
	"backend/controllers/response"
	"backend/helper"
	"backend/middleware"
	"backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

//login regis not found
//get user by id JWT

// LOGIN User "POST -> http://127.0.0.1:8080/login"
func LoginUserController(c echo.Context) error {
	user := models.User{}

	c.Bind(&user)
	password := user.Password

	// if err := config.DB.Where("PhoneNumber = ? AND password = ?", user.PhoneNumber, user.Password).First(&user).Error; err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }
	if err := config.DB.Table("user").Debug().Where("email = ? ", user.Email).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	matchPassword := matchPassword(user.Password, []byte(password))
	if !matchPassword {
		return c.JSON(http.StatusCreated, helper.BuildResponse("password salah", nil))
		// phoneNumber, _ := strconv.Atoi(user.PhoneNumber)
	}
	token, err := middleware.CreateToken(user.Id, "user")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil login",
		"User":    token,
	})

}

func matchPassword(hashedPassword string, password []byte) bool {
	byteHash := []byte(hashedPassword)
	if err := bcrypt.CompareHashAndPassword(byteHash, password); err != nil {
		return false
	}
	return true
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

	newreqeust := request.ReqUser{}
	c.Bind(&newreqeust)
	newuser := newreqeust.MapToDomain()

	newuser.Password = helper.CreateHash(newuser.Password)
	if err := config.DB.Table("user").Debug().Create(&newuser).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new user", response.MapToUser(newuser)))
}

//GET All User Data "GET -> http://127.0.0.1:8080/users"
func GetAllusercontrollers(c echo.Context) error {
	var users []models.User
	if err := config.DB.Table("user").Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all users", response.MapToBatchUser(users)))
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

	return c.JSON(http.StatusOK, helper.BuildResponse("success get user", response.MapToUser(user)))
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
	user.Password = helper.CreateHash(user.Password)

	return c.JSON(http.StatusOK, helper.BuildResponse("user deleted successfully", response.MapToUser(user)))
}

// EDIT Spesific User Data "PUT -> http://127.0.0.1:8080/users/:id"

func UpdateUserControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("uid"))
	fmt.Println(err)
	user := models.User{}

	if err := config.DB.Table("user").Debug().Where("id", id).Find(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "user not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	newreqeust := request.ReqUser{}

	c.Bind(&newreqeust)
	newuser := newreqeust.MapToDomain()

	if err := config.DB.Table("user").Debug().Where("id", id).Updates(&newuser).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	fmt.Println(user.Id)
	newuser.Id = user.Id
	newuser.Point = user.Point
	return c.JSON(http.StatusOK, helper.BuildResponse("success update user", response.MapToUser(newuser)))
}

func AddPointUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var user models.User
	var reqUser models.User
	if err := config.DB.First(&user, userId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusNotFound, "User not found")
	}
	if user.Id == 0 {
		return c.String(http.StatusNotFound, "User not found")
	}
	if err := c.Bind(&reqUser); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	user.Point = user.Point + reqUser.Point
	if err := config.DB.Save(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
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
