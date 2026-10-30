package dto

import (
	"strings"

	"github.com/codacube/go-microservice-banking/internal/errs"
)

type AccountRequest struct {
	CustomerID  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

type AccountResponse struct {
	AccountID string `json:"account_id"`
}

func (r AccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit at least 5000.00")
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "cheque" {
		return errs.NewValidationError("Account type should be saving or cheque")
	}
	return nil
}
