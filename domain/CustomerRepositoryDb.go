package domain

import (
	"database/sql"
	"time"

	"github.com/codacube/go-microservice-banking/errs"
	"github.com/codacube/go-microservice-banking/logger"
	_ "github.com/go-sql-driver/mysql"
)

// Server Side

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (c CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var rows *sql.Rows
	var err error

	/* Translating status is currently done in the service. However
	   if the adapter is something other than MySQL, then the expected
	   values might not work (because status in the DB could be
	   true/false, etc). The service doesn't know about adapter requirements,
	   and the adapter doesn't know about business requirements, etc.

	   Personally I would have a translation, or do this in the domain:

		if status == "active" {
			dbStatus = 1
		} elseif status == "inactive" {
			dbStatus = 0
		}
	*/

	if status == "" {
		findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		rows, err = c.client.Query(findAllSQL)
	} else {
		findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		rows, err = c.client.Query(findAllSQL, status)
	}

	if err != nil {
		logger.Error("Error whilst querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		if err := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Postcode, &customer.DateOfBirth, &customer.Status); err != nil {
			logger.Error("Error while getting all customers " + err.Error())
			return nil, errs.NewUnexpectedError("Error while scanning customers " + err.Error())
		}
		customers = append(customers, customer)
	}
	if err := rows.Err(); err != nil {
		logger.Error("Error while getting all customers " + err.Error())
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
			logger.Error("Customer not found")
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error while scanning customer " + err.Error())
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
