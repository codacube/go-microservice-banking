package service

import (
	"github.com/codacube/go-microservice-banking/internal/api/domain"
	"github.com/codacube/go-microservice-banking/internal/api/dto"
	"github.com/codacube/go-microservice-banking/internal/errs"
)

// CustomerService is the primary port that exposes customer use cases to
// inbound adapters, such as HTTP handlers.
type CustomerService interface {
	GetAllCustomers(status domain.CustomerStatus) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}

// DefaultCustomerService (Business Logic) implements the primary port and contains the
// application logic. It uses the CustomerRepository secondary port to access
// customer data without depending on a specific storage implementation.
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// Receiver function for DefaultCustomerService to implement the CustomerService interface
// (connects the primary port to the secondary port)
func (s DefaultCustomerService) GetAllCustomers(status domain.CustomerStatus) ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, c.ToDTO())
	}
	return response, err
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ByID(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDTO()

	return &response, nil
}

// Helper function to create a new instance of DefaultCustomerService with the provided repository
func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}
