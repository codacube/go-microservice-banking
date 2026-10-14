package service

import "github.com/codacube/go-microservice-banking/domain"

// Port (Primary) - Service
type CustomerService struct {
	GetAllCustomers func() ([]domain.Customer, error)
}

// Business Logic (accessing the secondary port - CustomerRepository)
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
