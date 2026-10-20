package domain

import (
	"github.com/codacube/go-microservice-banking/dto"
	"github.com/codacube/go-microservice-banking/errs"
)

// Business Side Layer - Domain Layer

type Customer struct {
	Id          string
	Name        string
	City        string
	Postcode    string
	DateOfBirth string
	Status      CustomerStatus
}

type CustomerStatus string

const (
	CustomerStatusActive   CustomerStatus = "active"
	CustomerStatusInactive CustomerStatus = "inactive"
)

func (c Customer) ToDTO() dto.CustomerResponse {
	response := dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Postcode:    c.Postcode,
		DateOfBirth: c.DateOfBirth,
		Status:      string(c.Status),
	}

	return response
}

// Ports (Secondary)
type CustomerRepository interface {
	FindAll(status CustomerStatus) ([]Customer, *errs.AppError)
	ByID(id string) (*Customer, *errs.AppError)
}
