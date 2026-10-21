package service

import (
	"time"

	"github.com/codacube/go-microservice-banking/domain"
	"github.com/codacube/go-microservice-banking/dto"
	"github.com/codacube/go-microservice-banking/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	// Validate incoming request
	err := req.Validate()
	if err != nil {
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
	response := newAccount.ToNewAccountResponseDTO()

	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
