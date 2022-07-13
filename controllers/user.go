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
	"strings"

	"github.com/go-playground/validator/v10"
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

	if err := config.DB.Table("users").Debug().Where("email = ?", user.Email).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	matchPassword := matchPassword(user.Password, []byte(password))
	if !matchPassword {
		return c.JSON(http.StatusCreated, helper.BuildResponse("password salah", nil))

	}
	token, err := middleware.CreateToken(user.Id, "user")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil login",
		"User Id": user.Id,
		"User":    token,
	})
}

func matchPassword(hashedPassword string, password []byte) bool {
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, password)
	return err == nil
}

func CreateUserControllers(c echo.Context) error {

	newrequest := request.ReqNewUser{}
	c.Bind(&newrequest)
	newuser := newrequest.MapToNewUser()
	validate := validator.New()
	if err := validate.Struct(newrequest); err != nil {
		var reasons []map[string]string
		invalids := err.(validator.ValidationErrors)
		for _, invalid := range invalids {
			reasons = append(reasons, map[string]string{invalid.Field(): strings.Split(invalid.Error(), "Error:")[1]})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var email string
	if err := config.DB.Table("users").Select("email").Where("email=? AND deleted_at is null", newuser.Email).Find(&email).Error; err != nil {
		return err
	}

	if email != "" {
		return c.String(http.StatusBadRequest, "email is already registered")
	}

	var phonenumber string
	if err := config.DB.Table("users").Select("phone_number").Where("phone_number=? AND deleted_at is null", newuser.PhoneNumber).Find(&phonenumber).Error; err != nil {
		return err
	}

	if phonenumber != "" {
		return c.String(http.StatusBadRequest, "phone number is already registered")
	}
	// newuser.Password = helper.CreateHash(newuser.Password)
	if err := config.DB.Table("users").Debug().Create(&newuser).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new user", response.MapToUser(newuser)))
}

//GET All User Data "GET -> http://127.0.0.1:8080/users"
func GetAllusercontrollers(c echo.Context) error {
	var users []models.User
	if err := config.DB.Table("users").Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all users", response.MapToBatchUser(users)))
}

// GET Spesific User Data "GET -> http://127.0.0.1:8080/users/:uid"
func GetUserControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	if err := config.DB.Table("users").First(&user, id).Error; err != nil {
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
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	if err := config.DB.Table("users").First(&user, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "user not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Table("users").Delete(&user).Error; err != nil {
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

	if err := config.DB.Table("users").Debug().Where("id", id).Find(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "user not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	newreqeust := request.ReqUser{}

	c.Bind(&newreqeust)
	newuser := newreqeust.MapToUser()

	if err := config.DB.Table("users").Debug().Where("id", id).Updates(&newuser).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	fmt.Println(user.Id)
	newuser.Id = user.Id
	newuser.Point = user.Point
	return c.JSON(http.StatusOK, helper.BuildResponse("success update user", response.MapToUser(newuser)))
}

func AddPointUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	pointuser := models.AddPointUser{}
	if err := c.Bind(&pointuser); err != nil {
		fmt.Println(err)
	}

	// userId, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return c.String(http.StatusBadRequest, "invalid id")
	// }
	var user models.User
	// var reqUser models.User
	if err := config.DB.First(&user, id).Error; err != nil {

		fmt.Println(err)
		return c.String(http.StatusNotFound, "User not found")
	}
	if user.Id == 0 {

		return c.String(http.StatusNotFound, "User not found")
	}
	fmt.Println(pointuser.Point)
	user.Point = user.Point + pointuser.Point
	fmt.Println(user.Point)
	if err := config.DB.Table("users").Debug().Where("id", id).Update("point", user.Point).Error; err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success add point user",
		"point":   user.Point,
	})

}
