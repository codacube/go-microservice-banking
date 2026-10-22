package dto

import (
	"strings"

	"github.com/codacube/go-microservice-banking/errs"
)

type NewTransactionRequest struct {
	AccountID string `json:"account_id"`
	// CustomerID      string  `json:"customer_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (r NewTransactionRequest) Validate() *errs.AppError {
	if strings.ToLower(r.TransactionType) != "widthdrawal" && strings.ToLower(r.TransactionType) != "deposit" {
		return errs.NewValidationError("A transaction must be of type 'withdrawal' or 'deposit'")
	}
	if r.Amount < 0 {
		return errs.NewValidationError("A transaction amount must be positive amount")
	}
	// TODO: Check that the amount is available in the account??
	return nil
}
