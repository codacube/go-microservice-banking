package domain

import "github.com/codacube/go-microservice-banking/errs"

// Business Side Layer - Domain Layer

type Customer struct {
	Id          string
	Name        string
	City        string
	Postcode    string
	DateOfBirth string
	Status      string
}

// Ports (Secondary)
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ByID(id string) (*Customer, *errs.AppError)
}
