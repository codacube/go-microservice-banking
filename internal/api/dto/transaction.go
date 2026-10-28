package dto

import (
	"strings"

	"github.com/codacube/go-microservice-banking/internal/errs"
)

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type NewTransactionRequest struct {
	AccountID       string  `json:"account_id"`
	CustomerID      string  `json:"customer_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

type NewTransactionResponse struct {
	TransactionID   string  `json:"transaction_id"`
	AccountID       string  `json:"account_id"`
	Balance         float64 `json:"new_balance"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (r NewTransactionRequest) Validate() *errs.AppError {
	if !r.IsTransactionTypeWithdrawal() && !r.IsTransactionTypeDeposit() {
		return errs.NewValidationError("A transaction must be of type 'withdrawal' or 'deposit'")
	}
	if r.Amount < 0 {
		return errs.NewValidationError("A transaction amount must be positive amount")
	}
	// TODO: Check that the amount is available in the account??
	return nil
}

func (r NewTransactionRequest) IsTransactionTypeWithdrawal() bool {
	return strings.ToLower(r.TransactionType) == WITHDRAWAL
}

func (r NewTransactionRequest) IsTransactionTypeDeposit() bool {
	return strings.ToLower(r.TransactionType) == DEPOSIT
}
