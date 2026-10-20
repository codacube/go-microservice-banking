package domain

import "github.com/codacube/go-microservice-banking/errs"

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningData string
	AccountType string
	Amount      float64
	Status      string
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
