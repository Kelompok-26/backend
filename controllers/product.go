package controllers

import (
	"backend/config"
	"backend/helper"
	"backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all products
func GetAllProductControllers(c echo.Context) error {
	var products []models.Product
	if err := config.DB.Table("product").Find(&products).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all products", products))
}

// get product by id
func GetProductControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
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

// create product by id
func CreateProductControllers(c echo.Context) error {
	product := models.Product{}
	c.Bind(&product)

	if err := config.DB.Table("product").Debug().Create(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new product", product))
}

// delete product by id
func DeleteProductControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product := models.Product{}
	if err := config.DB.Table("product").Where("id = ?", id).First(&product).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "product not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Table("product").Where("id = ?", id).Delete(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("product deleted successfully", product))
}

// update product by id
func UpdateProductControllers(c echo.Context) error {
	id := c.Param("id")
	product := models.Product{}

	if err := config.DB.Table("product").First(&product, "id = ?", id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "product not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	newproduct := models.Product{}
	c.Bind(&newproduct)
	fmt.Println("product", product)
	product.TypeProduct = newproduct.TypeProduct
	product.ProviderName = newproduct.ProviderName
	product.ProductName = newproduct.ProductName
	product.Nominal = newproduct.Nominal
	product.Stock = newproduct.Stock
	if err := config.DB.Table("product").Where("id = ?", id).Debug().Save(&product).Debug().Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success update product", product))
}
