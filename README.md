# Sample Microservices Go Project

Simple microservice project that showcases:

- Routing
- [Gorilla/mux](https://github.com/gorilla/mux)
- [Hexagonal Architecture](<https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)>)
- Connecting to a [MySQL](https://github.com/go-sql-driver/mysql) database
- Structured logging user [Uber Zap](https://github.com/uber-go/zap)

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

### Retrieving a Customer by ID

Acceptance Criteria

- [x] GET /customer/<id>
- [x] The URL should only accept numeric IDs
- [x] API should return customer as a JSON object and a 200 status code
- [x] In case the customer ID does not exist, the API should return http status code 404 (Not Found)
- [x] In case of an unexpected error, the API should return a status code 500 (Internal Server Error) along with the error message
- [x] Errors are returned as JSON objects
