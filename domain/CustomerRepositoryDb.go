<<<<<<< HEAD
package domain

import (
	"database/sql"
	"time"

	"github.com/codacube/go-microservice-banking/errs"
	"github.com/codacube/go-microservice-banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Server Side

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

type customerDBRecord struct {
	Id          string `db:"customer_id"`
	Name        string `db:"name"`
	City        string `db:"city"`
	Postcode    string `db:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}

func (db CustomerRepositoryDb) FindAll(status CustomerStatus) ([]Customer, *errs.AppError) {
	var err error
	records := make([]customerDBRecord, 0)

	if status == "" {
		findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = db.client.Select(&records, findAllSQL)
	} else {
		dbStatus, ok := databaseStatus(status)
		if !ok {
			logger.Error("Status must be active or inactive")
			return nil, errs.NewBadRequestError("Status must be active or inactive")
		}
		findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = db.client.Select(&records, findAllSQL, dbStatus)
	}

	if err != nil {
		logger.Error("Error whilst querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	customers := make([]Customer, 0, len(records))
	for _, record := range records {
		customer, ok := customerFromDBRecord(record)
		if !ok {
			logger.Error("Unexpected customer status in database")
			return nil, errs.NewUnexpectedError("Unexpected customer status in database")
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func databaseStatus(status CustomerStatus) (string, bool) {
	switch status {
	case CustomerStatusActive:
		return "1", true
	case CustomerStatusInactive:
		return "0", true
	default:
		return "", false
	}
}

func customerFromDBRecord(record customerDBRecord) (Customer, bool) {
	var status CustomerStatus
	switch record.Status {
	case "1":
		status = CustomerStatusActive
	case "0":
		status = CustomerStatusInactive
	default:
		return Customer{}, false
	}

	return Customer{
		Id:          record.Id,
		Name:        record.Name,
		City:        record.City,
		Postcode:    record.Postcode,
		DateOfBirth: record.DateOfBirth,
		Status:      status,
	}, true
}

func (db CustomerRepositoryDb) ByID(id string) (*Customer, *errs.AppError) {
	findByIDSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var record customerDBRecord
	err := db.client.Get(&record, findByIDSQL, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Customer not found")
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error while scanning customer " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	customer, ok := customerFromDBRecord(record)
	if !ok {
		logger.Error("Unexpected customer status in database")
		return nil, errs.NewUnexpectedError("Unexpected customer status in database")
	}

	return &customer, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	// The appended flags might be needed to for localhost development, but appears to work without it
	// sql.Open("mysql", "root:your_strong_password@tcp(localhost:3306)/banking?allowPublicKeyRetrieval=true&parseTime=true&columnsWithAlias=true")
	client, err := sqlx.Open("mysql", "root:your_strong_password@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client: client}
}
=======
package domain

import (
	"database/sql"
	"time"

	"github.com/codacube/go-microservice-banking/errs"
	"github.com/codacube/go-microservice-banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Server Side

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (db CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)

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
		err = db.client.Select(&customers, findAllSQL)
	} else {
		findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = db.client.Select(&customers, findAllSQL, status)
	}

	if err != nil {
		logger.Error("Error whilst querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (db CustomerRepositoryDb) ByID(id string) (*Customer, *errs.AppError) {
	findByIDSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var customer Customer
	err := db.client.Get(&customer, findByIDSQL, id)
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
	client, err := sqlx.Open("mysql", "root:your_strong_password@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client: client}
}
>>>>>>> a65c193 (SQLX changes applied)
