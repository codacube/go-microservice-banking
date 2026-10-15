package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (c CustomerRepositoryDb) FindAll() ([]Customer, error) {
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := c.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying customers table " + err.Error())
		return nil, err
		// panic(err)
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Postcode, &customer.DateOfBirth, &customer.Status)
		if err != nil {
			log.Println("Error while scanning customers " + err.Error())
			return nil, err
			// panic(err)
		}
		customers = append(customers, customer)
	}
	return customers, nil
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
