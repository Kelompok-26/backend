package controllers

import (
	"backend/config"
	"backend/controllers/request"
	"backend/controllers/response"
	"backend/helper"
	"backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllTransactionControllers(c echo.Context) error {
	var transaction []models.Transaction
	if err := config.DB.Table("transaction").Model(&transaction).Debug().Preload("User").Preload("Product").Find(&transaction).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all transaction", response.MapToBatchTransaction(transaction)))
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

	return c.JSON(http.StatusOK, helper.BuildResponse("success get transaction", response.MapToBatchTransaction(transaction)))
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

	return c.JSON(http.StatusOK, helper.BuildResponse("success get transaction", response.MapToBatchTransaction(transaction)))
}

// delete transaction by id
func DeleteTansactionControllers(c echo.Context) error {
	id := c.Param("id")
	transaction := models.Transaction{}

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

	return c.JSON(http.StatusOK, helper.BuildResponse("transaction deleted successfully", transaction.Id))
}

// update by id
func UpdatetransactionControllers(c echo.Context) error {
	id := c.Param("id")
	transaction := models.Transaction{}

	if err := config.DB.Table("transaction").Where("id", id).Find(&transaction).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "transaction not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	newtransaction := models.Transaction{}
	c.Bind(&newtransaction)

	transaction.User = newtransaction.User
	transaction.Product = newtransaction.Product
	// transaction.Point = newtransaction.Point
	// transaction.Total = newtransaction.Total
	if err := config.DB.Table("transaction").Where("id", id).Save(&transaction).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success update transaction", response.MapToTransaction(transaction)))
}

func UserCreateTransactionsController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	reqtrans := request.ReqTransaction{}

	if err := c.Bind(&reqtrans); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error ")
	}
	transactions := reqtrans.MapToTransaction()
	transactions.UserId = id
	// Get User By ID
	var user models.User
	if err := config.DB.Table("users").First(&user, id).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusNotFound, "User not found")
	}
	// Get product By ID
	var product models.Product
	if err := config.DB.Table("products").First(&product, transactions.ProductId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusNotFound, "product not found")
	}
	// Point User < product.Point
	if user.Point < product.Point {
		fmt.Println(user.Point)
		return c.String(http.StatusBadRequest, "Not Enough Point")
	}

	// Stock = 0
	if product.Stock == 0 {
		return c.String(http.StatusBadRequest, "Stock Out")
	} else {
		// Reduce product Stock By 1
		if err := config.DB.Table("products").Model(&product).Update("stock", product.Stock-1).Error; err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "Internal Server Error ")
		}
	}
	// Reduce Point User
	if err := config.DB.Table("users").Model(&user).Where("id = ?", user.Id).Update("point", user.Point-product.Point).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if err := config.DB.Table("transaction").Create(&transactions).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	if err := config.DB.Table("transaction").Preload("User").Preload("Product").Find(&transactions).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, response.MapToTransaction(transactions))
}
