package controllers

import (
	"backend/config"
	"backend/controllers/response"
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

	if err := config.DB.Table("products").Debug().Create(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new product", response.MapToProduct(product)))
}

// GET All Product Data "GET ->http://127.0.0.1:8080/products"
func GetAllProductControllers(c echo.Context) error {
	var products []models.Product
	if err := config.DB.Table("products").Find(&products).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all products", response.MapToBatchProduct(products)))
}

// GET Product Paket Data "GET ->http://127.0.0.1:8080/products/PaketData"
func GetPaketData(c echo.Context) error {
	product := models.Product{}

	if err := config.DB.Where("type_product", "Paket Data").First(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, helper.BuildResponse("success get all Paket Data", response.MapToProduct(product)))
}

// GET Product Pulsa "GET ->http://127.0.0.1:8080/products/Pulsa"
func GetPulsa(c echo.Context) error {
	product := models.Product{}

	if err := config.DB.Where("type_product", "Pulsa").First(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, helper.BuildResponse("success get all Pulsa", response.MapToProduct(product)))
}

// GET Product Emoney "GET ->http://127.0.0.1:8080/products/Emoney"
func GetEmoney(c echo.Context) error {
	product := models.Product{}

	if err := config.DB.Where("type_product", "E-Money").First(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, helper.BuildResponse("success get all E-money", response.MapToProduct(product)))
}

// GET Product Cashout "GET ->http://127.0.0.1:8080/products/Cashout"
func GetCashout(c echo.Context) error {
	product := models.Product{}

	if err := config.DB.Where("type_product", "Cashout").First(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, helper.BuildResponse("success get all Cashout", response.MapToProduct(product)))
}

// GET Spesific Product Data using ID "PUT -> http://127.0.0.1:8080/products/:pid"
func GetProductControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("pid"))
	product := models.Product{}
	if err := config.DB.Table("products").First(&product, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "product not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get product", response.MapToProduct(product)))
}

// EDIT Product "PUT -> http://127.0.0.1:8080/products/:pid"
func UpdateProductControllers(c echo.Context) error {
	id := c.Param("id")
	product := models.Product{}

	if err := config.DB.Table("products").Where("id", id).Find(&product).Error; err != nil {
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

	if err := config.DB.Table("products").Where("id", id).Save(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success update redeem", response.MapToProduct(product)))
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
	if err := config.DB.Table("products").Where("pid = ?", id).First(&product).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "product not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Table("products").Where("pid = ?", id).Delete(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("product deleted successfully", response.MapToProduct(product)))
}
