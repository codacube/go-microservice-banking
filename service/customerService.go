package service

import "github.com/codacube/go-microservice-banking/domain"

// CustomerService is the primary port that exposes customer use cases to
// inbound adapters, such as HTTP handlers.
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

// DefaultCustomerService implements the primary port and contains the
// application logic. It uses the CustomerRepository secondary port to access
// customer data without depending on a specific storage implementation.
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// Receiver function for DefaultCustomerService to implement the CustomerService interface
func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// Helper function to create a new instance of DefaultCustomerService with the provided repository
func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}
