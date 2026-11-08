package service

import (
	"testing"

	realdomain "github.com/codacube/go-microservice-banking/internal/api/domain"
	"github.com/codacube/go-microservice-banking/internal/api/dto"
	"github.com/codacube/go-microservice-banking/internal/api/mocks/domain"
	"github.com/codacube/go-microservice-banking/internal/errs"
	"go.uber.org/mock/gomock"
)

func setup(t *testing.T) (*domain.MockAccountRepository, AccountService) {
	t.Helper()
	ctrl := gomock.NewController(t)
	mockRepo := domain.NewMockAccountRepository(ctrl)
	service := NewAccountService(mockRepo)

	return mockRepo, service
}

func Test_should_return_a_validation_error_response_if_the_request_is_invalid(t *testing.T) {
	// Arrange
	request := dto.AccountRequest{CustomerID: "100", AccountType: "Saving", Amount: 0}
	service := NewAccountService(nil)

	// Act
	_, appError := service.NewAccount(request)

	// Assert
	if appError == nil {
		t.Error("Expected an error but got nil")
	}
}

func Test_should_return_an_error_from_the_server_side_if_the_account_cannot_be_created(t *testing.T) {
	// Arrange
	mockRepo, service := setup(t)
	request := dto.AccountRequest{CustomerID: "100", AccountType: "Saving", Amount: 6000}
	account := realdomain.NewAccount(request.CustomerID, request.AccountType, request.Amount)
	mockRepo.EXPECT().Save(account).Return(nil, errs.NewUnexpectedError("Database error"))

	// Act
	_, appError := service.NewAccount(request)

	// Assert
	if appError == nil {
		t.Error("Expected an error but got nil")
	}
}

func Test_should_return_an_account_response_if_the_account_is_created_successfully(t *testing.T) {
	// Arrange
	mockRepo, service := setup(t)
	request := dto.AccountRequest{CustomerID: "100", AccountType: "Saving", Amount: 6000}
	account := realdomain.NewAccount(request.CustomerID, request.AccountType, request.Amount)

	accountWithID := account
	accountWithID.AccountID = "201"
	mockRepo.EXPECT().Save(account).Return(&accountWithID, nil)

	// Act
	newAccount, appError := service.NewAccount(request)

	// Assert
	if appError != nil {
		t.Errorf("Expected no error but got %v", appError)
	}
	if newAccount.AccountID != "201" {
		t.Errorf("Expected account ID '201' but got %s", newAccount.AccountID)
	}
}
