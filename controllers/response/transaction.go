package response

import (
	"backend/helper"
	"backend/models"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Point int    `json:"point"`
}

type Product struct {
	Id           int    `json:"id"`
	TypeProduct  string `json:"type_product"`
	ProviderName string `json:"provider_name"`
	ProductName  string `json:"product_name"`
	Nominal      int    `json:"nominal"`
	Point        int    `json:"point"`
	Stock        int    `json:"stock"`
}

type Transaction struct {
	Id        int     `json:"id"`
	User      User    `json:"user"`
	Number    string  `json:"number"`
	Product   Product `json:"product"`
	CreatedAt string  `json:"created_at"`
}

func MapToTransaction(transaction models.Transaction) Transaction {
	return Transaction{
		Id:     transaction.Id,
		Number: transaction.Number,
		User: User{
			Id:    transaction.User.Id,
			Name:  transaction.User.Name,
			Email: transaction.User.Email,
			Point: transaction.User.Point},
		Product: Product{
			Id:           transaction.Product.Id,
			TypeProduct:  transaction.Product.TypeProduct,
			ProviderName: transaction.Product.ProviderName,
			ProductName:  transaction.Product.ProductName,
			Nominal:      transaction.Product.Nominal,
			Point:        transaction.Product.Point,
			Stock:        transaction.Product.Stock,
		},
		CreatedAt: helper.ConvertDatetimeToString(transaction.CreatedAt),
	}
}

func MapToBatchTransaction(transactions []models.Transaction) []Transaction {
	var responses []Transaction

	for _, transaction := range transactions {
		responses = append(responses, MapToTransaction(transaction))
	}
	return responses
}
