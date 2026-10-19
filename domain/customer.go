package domain

import (
	"github.com/codacube/go-microservice-banking/dto"
	"github.com/codacube/go-microservice-banking/errs"
)

// Business Side Layer - Domain Layer

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string `db:"full_name"`
	City        string `db:"city"`
	Postcode    string `db:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}

	return statusAsText
}

func (c Customer) ToDTO() dto.CustomerResponse {
	response := dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Postcode:    c.Postcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}

	return response
}

// Ports (Secondary)
type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ByID(id string) (*Customer, *errs.AppError)
}
