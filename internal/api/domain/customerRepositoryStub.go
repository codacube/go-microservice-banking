package domain

import "github.com/codacube/go-microservice-banking/internal/errs"

// Adapters
type CustomerRepositoryStub struct {
	customers []Customer
}

// Go supports Duck Typing, so we don't need to explicitly declare that CustomerRepositoryStub implements CustomerRepository interface. If it has the same methods, it is considered to implement the interface.
func (c *CustomerRepositoryStub) FindAll(status CustomerStatus) ([]Customer, *errs.AppError) {
	if status == "" {
		return c.customers, nil
	}
	if status != CustomerStatusActive && status != CustomerStatusInactive {
		return nil, errs.NewBadRequestError("Status must be active or inactive")
	}

	customers := make([]Customer, 0)
	for _, customer := range c.customers {
		if customer.Status == status {
			customers = append(customers, customer)
		}
	}
	return customers, nil
}

func (c *CustomerRepositoryStub) ByID(id string) (*Customer, *errs.AppError) {
	for _, customer := range c.customers {
		if customer.Id == id {
			return &customer, nil
		}
	}
	return nil, errs.NewNotFoundError("Customer not found")
}

func NewCustomerRepositoryStub() *CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", Name: "John", City: "Auckland", Postcode: "10001", DateOfBirth: "1990-01-01", Status: CustomerStatusActive},
		{Id: "2", Name: "Jane", City: "Wellington", Postcode: "90001", DateOfBirth: "1992-02-02", Status: CustomerStatusInactive},
		{Id: "3", Name: "Bob", City: "Christchurch", Postcode: "60601", DateOfBirth: "1994-03-03", Status: CustomerStatusActive},
	}
	return &CustomerRepositoryStub{customers: customers}
}
