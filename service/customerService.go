package service

import (
	"github.com/codacube/go-microservice-banking/domain"
	"github.com/codacube/go-microservice-banking/errs"
)

// CustomerService is the primary port that exposes customer use cases to
// inbound adapters, such as HTTP handlers.
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, *errs.AppError)
	GetCustomer(id string) (*domain.Customer, *errs.AppError)
}

// DefaultCustomerService (Business Logic) implements the primary port and contains the
// application logic. It uses the CustomerRepository secondary port to access
// customer data without depending on a specific storage implementation.
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// Receiver function for DefaultCustomerService to implement the CustomerService interface
// (connects the primary port to the secondary port)
func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ByID(id)
}

// Helper function to create a new instance of DefaultCustomerService with the provided repository
func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}
