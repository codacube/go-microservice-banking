package domain

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
}
