package domain

import (
	"database/sql"
	"log"
	"time"

	"github.com/codacube/go-microservice-banking/errs"
	_ "github.com/go-sql-driver/mysql"
)

// Server Side

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (c CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {
	findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := c.client.Query(findAllSQL)
	if err != nil {
		log.Println("Error whilst querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		if err := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Postcode, &customer.DateOfBirth, &customer.Status); err != nil {
			log.Println("Error while getting all customers " + err.Error())
			return nil, errs.NewUnexpectedError("Error while scanning customers " + err.Error())
		}
		customers = append(customers, customer)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error while getting all customers " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return customers, nil
}

func (c CustomerRepositoryDb) ByID(id string) (*Customer, *errs.AppError) {
	findByIDSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := c.client.QueryRow(findByIDSQL, id)

	var customer Customer
	err := row.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Postcode, &customer.DateOfBirth, &customer.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		log.Println("Error while scanning customer " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &customer, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	// The appended flags might be needed to for localhost development, but appears to work without it
	// sql.Open("mysql", "root:your_strong_password@tcp(localhost:3306)/banking?allowPublicKeyRetrieval=true&parseTime=true")
	client, err := sql.Open("mysql", "root:your_strong_password@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client: client}
}
