package controllers

import (
	"backend/config"
	"backend/helper"
	"backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Create New Product "POST -> http://127.0.0.1:8080/products"
// {
// 		"TypeProduct": "",
// 		"ProviderName": "",
//		"ProductName": "",
// 		"Nominal": ,
//		"Point": ,
// 		"Stock":
// }
func CreateProductControllers(c echo.Context) error {
	product := models.Product{}
	c.Bind(&product)

	if err := config.DB.Table("product").Debug().Create(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new product", product))
}

// GET All Product Data "GET ->http://127.0.0.1:8080/products"
func GetAllProductControllers(c echo.Context) error {
	var products []models.Product
	if err := config.DB.Table("product").Find(&products).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all products", products))
}

// GET Product Paket Data "GET ->http://127.0.0.1:8080/products/PaketData"
func GetPaketData(c echo.Context) error {
	product := models.Product{}

	if err := config.DB.Where(product.TypeProduct, "Paket Data").First(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, helper.BuildResponse("success get all Paket Data", product))
}

// GET Product Pulsa "GET ->http://127.0.0.1:8080/products/Pulsa"
func GetPulsa(c echo.Context) error {
	product := models.Product{}

	if err := config.DB.Where(product.TypeProduct, "Pulsa").First(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, helper.BuildResponse("success get all Pulsa", product))
}

// GET Product Emoney "GET ->http://127.0.0.1:8080/products/Emoney"
func GetEmoney(c echo.Context) error {
	product := models.Product{}

	if err := config.DB.Where(product.TypeProduct, "Emoney").First(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, helper.BuildResponse("success get all E-money", product))
}

// GET Product Cashout "GET ->http://127.0.0.1:8080/products/Cashout"
func GetCashout(c echo.Context) error {
	product := models.Product{}

	if err := config.DB.Where(product.TypeProduct, "Cashout").First(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, helper.BuildResponse("success get all Cashout", product))
}

// GET Spesific Product Data using ID "PUT -> http://127.0.0.1:8080/products/:pid"
func GetProductControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("pid"))
	product := models.Product{}
	if err := config.DB.Table("product").First(&product, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "product not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get product", product))
}

// EDIT Product "PUT -> http://127.0.0.1:8080/products/:pid"
func UpdateProductControllers(c echo.Context) error {
	id := c.Param("id")
	product := models.Product{}

	if err := config.DB.Table("product").Where("id", id).Find(&product).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "product not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	newproduct := models.Product{}
	c.Bind(&newproduct)
	product.TypeProduct = newproduct.TypeProduct
	product.ProviderName = newproduct.ProviderName
	product.ProductName = newproduct.ProductName
	product.Point = newproduct.Point
	product.Nominal = newproduct.Nominal
	product.Stock = newproduct.Stock

	if err := config.DB.Table("product").Where("id", id).Save(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success update redeem", product))
}

// newproduct := models.Product{}
// c.Bind(&newproduct)
// product.TypeProduct = newproduct.TypeProduct
// product.ProviderName = newproduct.ProviderName
// product.ProductName = newproduct.ProductName
// product.Point = newproduct.Point
// product.Nominal = newproduct.Nominal
// product.Stock = newproduct.Stock
// DELETE Product "DELETE -> http://127.0.0.1:8080/products/:pid"
func DeleteProductControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("pid"))
	product := models.Product{}
	if err := config.DB.Table("product").Where("pid = ?", id).First(&product).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "product not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Table("product").Where("pid = ?", id).Delete(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("product deleted successfully", product))
}
