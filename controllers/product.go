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
	var productId int
	if err := config.DB.Table("products").Select("id").
	Where("product_name=? AND type_product=? AND provider_name=? AND deleted_at is null", 
	product.ProductName, product.TypeProduct, product.ProviderName).
	Find(&productId).Error; err != nil {
		return err
	}
	if productId != 0 {
		return c.String(http.StatusBadRequest, "already registered")
	}
	if err := config.DB.Table("products").Debug().Create(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if product.ProductName == "" {
		return c.String(http.StatusBadRequest, "product name is nil")
	}
	if product.ProviderName == "" {
		return c.String(http.StatusBadRequest, "provider name is nil")
	}

	if product.TypeProduct == "" {
		return c.String(http.StatusBadRequest, "type product is nil")
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
	var product []models.Product

	if err := config.DB.Where("type_product", "Paket Data").Find(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, helper.BuildResponse("success get all Paket Data", response.MapToBatchProduct(product)))
}

// GET Product Pulsa "GET ->http://127.0.0.1:8080/products/Pulsa"
func GetPulsa(c echo.Context) error {

	var product []models.Product

	if err := config.DB.Where("type_product", "Pulsa").Find(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, helper.BuildResponse("success get all Pulsa", response.MapToBatchProduct(product)))
}

// GET Product Emoney "GET ->http://127.0.0.1:8080/products/Emoney"
func GetEmoney(c echo.Context) error {
	var product []models.Product

	if err := config.DB.Where("type_product", "E-Money").Find(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, helper.BuildResponse("success get all E-money", response.MapToBatchProduct(product)))
}

// GET Product Cashout "GET ->http://127.0.0.1:8080/products/Cashout"
func GetCashout(c echo.Context) error {
	var product []models.Product

	if err := config.DB.Where("type_product", "Cashout").Find(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, helper.BuildResponse("success get all Cashout", response.MapToBatchProduct(product)))
}

// GET Spesific Product Data using ID "PUT -> http://127.0.0.1:8080/products/:pid"
func GetProductControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product := models.Product{}
	if err := config.DB.Table("products").Find(&product, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "product not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if product.Id == 0 {
		return c.String(http.StatusNotFound, "product not found")
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get product", response.MapToProduct(product)))
}

// EDIT Product "PUT -> http://127.0.0.1:8080/products/:pid"
func UpdateProductControllers(c echo.Context) error {
	id := c.Param("id")
	product := models.Product{}


	if err := config.DB.Table("products").Debug().Where("id", id).Find(&product).Error; err != nil {
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
	
	var productId int
	if err := config.DB.Table("products").Select("id").
	Where("product_name=? AND type_product=? AND provider_name=? AND deleted_at is null", 
	product.ProductName, product.TypeProduct, product.ProviderName).
	Find(&productId).Error; err != nil {
		return err
	}
	if productId != 0 {
		return c.String(http.StatusBadRequest, "already registered")
	}
	
	if err := config.DB.Table("products").Debug().Where("id", id).Updates(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if product.Id == 0 {
		return c.String(http.StatusNotFound, "product not found")
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success update product", response.MapToProduct(product)))
}

// DELETE Product "DELETE -> http://127.0.0.1:8080/products/:pid"
func DeleteProductControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("pid"))
	product := models.Product{}
	if err := config.DB.Table("products").Debug().Where("id = ?", id).First(&product).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "product not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Table("products").Debug().Where("id = ?", id).Delete(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("product deleted successfully", response.MapToProduct(product)))
}
