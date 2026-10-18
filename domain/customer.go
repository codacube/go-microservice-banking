package domain

import "github.com/codacube/go-microservice-banking/errs"

// Business Side Layer - Domain Layer

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Postcode    string `db:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

// Ports (Secondary)
type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ByID(id string) (*Customer, *errs.AppError)
}
