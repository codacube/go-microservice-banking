package domain

import (
	"github.com/codacube/go-microservice-banking/internal/api/dto"
)

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	TransactionID   string  `db:"transaction_id"`
	AccountID       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

func (t Transaction) IsWithdrawal() bool {
	return t.TransactionType == WITHDRAWAL
}

func (t Transaction) ToDTO() dto.TransactionResponse {
	return dto.TransactionResponse{TransactionID: t.TransactionID, AccountID: t.AccountID, Balance: t.Amount, TransactionType: t.TransactionType, TransactionDate: t.TransactionDate}
}
