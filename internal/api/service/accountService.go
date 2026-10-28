package service

import (
	"time"

	"github.com/codacube/go-microservice-banking/internal/api/domain"
	"github.com/codacube/go-microservice-banking/internal/api/dto"
	"github.com/codacube/go-microservice-banking/internal/errs"
)

type AccountService interface {
	NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	NewTransaction(request dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	// Validate incoming request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Create a new account
	a := domain.Account{
		AccountID:   "",
		CustomerID:  req.CustomerID,
		OpeningDate: time.Now().Format("2006-01-02T15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToDTO()

	return &response, nil
}

func (s DefaultAccountService) NewTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	// Validate incoming request
	if err := req.Validate(); err != nil {
		return nil, err
	}
	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindBy(req.AccountID)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}

	// Create a new transaction
	t := domain.Transaction{
		TransactionID:   "",
		AccountID:       req.AccountID,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02T15:04:05"),
	}
	transaction, err := s.repo.SaveTransaction(t)
	if err != nil {
		return nil, err
	}
	response := transaction.ToDTO()

	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
