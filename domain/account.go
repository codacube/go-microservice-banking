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

func (a Account) ToDTO() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountID: a.AccountID}
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount >= amount
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
	FindBy(accountID string) (*Account, *errs.AppError)
}
