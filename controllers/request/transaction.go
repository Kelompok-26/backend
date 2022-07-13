package request

import (
	"backend/models"
)

type ReqTransaction struct {
	Number    string `json:"number"`
	ProductId int    `json:"product_id"`
}

func (transaction *ReqTransaction) MapToTransaction() models.Transaction {
	return models.Transaction{
		ProductId: transaction.ProductId,
		Number: transaction.Number,
	}
}
