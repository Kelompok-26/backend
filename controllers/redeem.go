package controllers

import (
	"backend/config"
	"backend/helper"
	"backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllredeemcontrollers(c echo.Context) error {
	var redeem []models.Redeem
	if err := config.DB.Model(&redeem).Debug().Preload("User").Find(&redeem).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all redeem", redeem))
}

func GetredeemByUserIDcontrollers(c echo.Context) error {
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

func CreateRedeemscontrollers(c echo.Context) error {
	redeem := models.InsertRedeem{}
	c.Bind(&redeem)

	if err := config.DB.Table("redeem").Create(&redeem).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new redeem", redeem))
}
