package domain

import "github.com/codacube/go-microservice-banking/errs"

// Adapters
type CustomerRepositoryStub struct {
	customers []Customer
}

// Go supports Duck Typing, so we don't need to explicitly declare that CustomerRepositoryStub implements CustomerRepository interface. If it has the same methods, it is considered to implement the interface.
func (c *CustomerRepositoryStub) FindAll() ([]Customer, *errs.AppError) {
	return c.customers, nil
}

func NewCustomerRepositoryStub() *CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", Name: "John", City: "Auckland", Postcode: "10001", DateOfBirth: "1990-01-01", Status: "1"},
		{Id: "2", Name: "Jane", City: "Wellington", Postcode: "90001", DateOfBirth: "1992-02-02", Status: "0"},
		{Id: "3", Name: "Bob", City: "Christchurch", Postcode: "60601", DateOfBirth: "1994-03-03", Status: "1"},
	}
	return &CustomerRepositoryStub{customers: customers}
}
