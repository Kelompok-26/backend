package response

import (
	"backend/helper"
	"backend/models"
)

type ResProduct struct {
	Id           int    `json:"id"`
	TypeProduct  string `json:"type_product"`
	ProviderName string `json:"provider_name"`
	ProductName  string `json:"product_name"`
	Nominal      int    `json:"nominal"`
	Point        int    `json:"point"`
	Stock        int    `json:"stock"`
	CreatedAt    string `json:"created_at"`
}

func MapToProduct(product models.Product) ResProduct {
	return ResProduct{
		Id:           product.Id,
		TypeProduct:  product.TypeProduct,
		ProviderName: product.ProviderName,
		ProductName:  product.ProductName,
		Nominal:      product.Nominal,
		Point:        product.Point,
		Stock:        product.Stock,
		CreatedAt:    helper.ConvertDatetimeToString(product.CreatedAt),
	}
}

func MapToBatchProduct(products []models.Product) []ResProduct {
	var responses []ResProduct

	for _, product := range products {
		responses = append(responses, MapToProduct(product))
	}
	return responses
}
