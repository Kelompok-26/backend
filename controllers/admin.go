package controllers

import (
	"backend/config"
	"backend/middleware"
	"backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginAdminController(c echo.Context) error {
	admin := models.Admin{}
	c.Bind(&admin)

	if err := config.DB.Table("admin").Where("phone_number = ? AND password = ?", admin.PhoneNumber, admin.Password).First(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	//
	token, err := middleware.CreateToken(admin.ID, string(admin.PhoneNumber), "admin")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil login",
		"admin":   token,
	})
}
