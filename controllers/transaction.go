package controllers

import (
	"backend/config"
	"backend/helper"
	"backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllTransactionControllers(c echo.Context) error {
	var transaction []models.Transaction
	if err := config.DB.Table("transaction").Model(&transaction).Debug().Preload("User").Preload("Product").Find(&transaction).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all transaction", transaction))
}
func GetTransactionByIdControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction := []models.Transaction{}
	if err := config.DB.Table("transaction").Model(&transaction).Debug().Where("id", id).Preload("User").Preload("Product").Find(&transaction).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "transaction not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get transaction", transaction))
}

func GetTransactionByIdUserControllers(c echo.Context) error {
	user_id, _ := strconv.Atoi(c.Param("user_id"))
	transaction := []models.Transaction{}
	if err := config.DB.Model(&transaction).Debug().Where("user_id", user_id).Preload("User").Preload("Product").Find(&transaction).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "redeem not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get transaction", transaction))
}

func CreateTransactionsControllers(c echo.Context) error {
	transaction := models.InsertTransaction{}
	c.Bind(&transaction)

	if err := config.DB.Table("transaction").Create(&transaction).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new transaction", transaction))
}

// delete transaction by id
func DeleteTansactionControllers(c echo.Context) error {
	id := c.Param("id")
	transaction := models.InsertTransaction{}

	if err := config.DB.Table("transaction").Where("id", id).Find(&transaction).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "transaction not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Table("transaction").Where("id", id).Delete(&transaction).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("transaction deleted successfully", transaction))
}

// update by id
func UpdatetransactionControllers(c echo.Context) error {
	id := c.Param("id")
	transaction := models.InsertTransaction{}

	if err := config.DB.Table("transaction").Where("id", id).Find(&transaction).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "transaction not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	newtransaction := models.InsertTransaction{}
	c.Bind(&newtransaction)

	transaction.User = newtransaction.User
	transaction.Product = newtransaction.Product
	transaction.Point = newtransaction.Point
	transaction.Total = newtransaction.Total
	if err := config.DB.Table("transaction").Where("id", id).Save(&transaction).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success update transaction", transaction))
}
