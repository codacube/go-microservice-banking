package domain

import (
	"github.com/codacube/go-microservice-banking/dto"
)

type Transaction struct {
	TransactionID   string
	AccountID       string
	Amount          float64
	TransactionType string
	TransactionDate string
}

func (t Transaction) ToNewTransactionResponseDTO() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{TransactionID: t.TransactionID, Amount: t.Amount}
}
