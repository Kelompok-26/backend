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

	if err := config.DB.Where("nomor_hp = ? AND password = ?", admin.PhoneNumber, admin.Password).First(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	//
	token, err := middleware.CreateToken(admin.ID, admin.PhoneNumber)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil login",
		"admin":   token,
	})
}


