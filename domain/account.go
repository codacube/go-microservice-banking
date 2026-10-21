package domain

import (
	"github.com/codacube/go-microservice-banking/dto"
	"github.com/codacube/go-microservice-banking/errs"
)

type Account struct {
	AccountID   string
	CustomerID  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDTO() dto.NewAccountResponse {
	return dto.NewAccountResponse{a.AccountID}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
