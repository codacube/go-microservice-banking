# Sample Microservices Go Project

Simple microservice project that showcases:

- Routing
- GET/POST
- [Gorilla/mux](https://github.com/gorilla/mux)
- [Hexagonal Architecture](<https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)>)
- Connecting to a [MySQL](https://github.com/go-sql-driver/mysql) database
- Structured logging user [Uber Zap](https://github.com/uber-go/zap)
- [SQLX](https://github.com/jmoiron/sqlx)
- Data Transfer Objects
- Environment variables

# Architecture / Structure

- **app/** handles HTTP
- **service/** co-ordinates use cases
- **domain/** holds business types nd repository interfaces
- **dto/** defines API shapes

## For testing

- GUI - Use [Bruno](https://www.usebruno.com/)
- CLI - curl

```bash
curl -i http://localhost:8000/greet
curl -i -H "Content-Type: application/xml" http://localhost:8000/customers
curl -i -H "Content-Type: application/json" http://localhost:8000/customers
```

## User Stories

### Retrieving all Customers

Acceptance Criteria

- [x] GET /customers
- [x] API should only return a JSON response
- [x] In case of an unexpected error, the API should return a status code 500 (Internal Server Error) along with the error message (as JSON object)
- [x] API should return HTTP status code 200 and the data upon successfully retrieving all customers from the server side
- [x] The API should provide an option to fetch the customers by status, eg:
  - [x] /customers?status=active - return all active customers
  - [x] /customers?status=inactive - return all inactive customers
  - [x] should return all customers
  - [x] unsupported status values should return HTTP status code 400 (Bad Request)

### Retrieving a Customer by ID

Acceptance Criteria

- [x] GET /customer/<id>
- [x] The URL should only accept numeric IDs
- [x] API should return customer as a JSON object and a 200 status code
- [x] In case the customer ID does not exist, the API should return http status code 404 (Not Found)
- [x] In case of an unexpected error, the API should return a status code 500 (Internal Server Error) along with the error message
- [x] Errors are returned as JSON objects

### Write an API to create a new accoutn for an existing customer

Acceptance Criteria

- [x] A new account can only be opened with a minimum deposit of 5000.00
- [x] Accounts can only be savings or cheque
- [x] In case of an unexpected error, the API should return a status code 500 (Internal Server Error) along with the error message
- [x] The API should return the new account ID, when the new account is opened with the status code as 201 (CREATED)

### Create a transaction in a bank account

Acceptance Criteria

Write an API to create a new transaction for an existing customer

- [x] Transaction can only be "withdrawal" or "deposit"
- [x] Amount cannot be negative
- [x] Withdrawal amount should be available in the account
- [x] Successful transaction should respond with the updated balance and the transaction id
- [x] Error handling should be done for a bad request, validation and unexpected errors from the server side and should return the appropriate http status code with a message

# AI Recommendations

The main architectural improvement I’d consider is making the layer boundaries more consistent:

1. **Keep the domain independent of API DTOs.** `Customer` and `Account` currently have `ToDTO` methods, so domain types import the API-facing `dto` package. Move that mapping to the service or HTTP boundary if you want the domain to remain independent of the API representation.
2. **Keep transport concerns out of domain behavior.** Request validation currently lives on DTOs in `account.go` and `transaction.go`. Basic shape/input checks can happen at the HTTP boundary; rules that express business constraints belong in the service or domain.
3. **Choose a consistent persistence mapping style.** `CustomerRepositoryDb.go` explicitly converts a database record into a domain type, while `AccountRepositoryDb.go` scans directly into `Account`. Either can work; consistency would make the persistence boundary easier to reason about.

I also noticed a couple of **specific correctness issues** worth prioritizing over a structural rewrite:

- In `AccountRepositoryDb.go`, the transaction insert error is discarded. That error should abort and roll back the operation. Separately, checking funds before the database transaction can allow concurrent withdrawals to overdraw an account; enforce the balance condition atomically in the database.

I wouldn’t introduce more layers or reorganize everything yet. Keep the current shape, fix the transactional and validation issues, and gradually tighten the boundaries as the service grows.
