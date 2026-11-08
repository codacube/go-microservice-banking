package domain

import (
	"time"

	"github.com/codacube/go-microservice-banking/internal/api/dto"
	"github.com/codacube/go-microservice-banking/internal/errs"
)

type Account struct {
	AccountID   string
	CustomerID  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

const dbTSLayout = "2006-01-02T15:04:05"

func (a Account) ToDTO() *dto.AccountResponse {
	return &dto.AccountResponse{AccountID: a.AccountID}
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount >= amount
}

//go:generate go tool go.uber.org/mock/mockgen -source=account.go -destination=../mocks/domain/mock_account_repository.go -package=domain
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
	FindBy(accountID string) (*Account, *errs.AppError)
}

func NewAccount(customerID string, accountType string, amount float64) Account {
	return Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-01-02T15:04:05"),
		AccountType: accountType,
		Amount:      amount,
		Status:      "1",
	}
}
