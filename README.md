# Sample Microservices Go Project

Simple microservice project that showcases:

- Routing
- [Gorilla/mux](https://github.com/gorilla/mux)
- [Hexagonal Architecture](<https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)>)
- Connecting to a [MySQL](https://github.com/go-sql-driver/mysql) database

## For testing

- GUI - Use [Bruno](https://www.usebruno.com/)
- CLI - curl

```bash
curl -i http://localhost:8000/greet
curl -i -H "Content-Type: application/xml" http://localhost:8000/customers
curl -i -H "Content-Type: application/json" http://localhost:8000/customers
```

## User Story for Retrieving a Customer by ID

Acceptance Criteria

- [x] The URL should only accept numeric IDs
- [x] API should return customer as a JSON object
- [x] In case the customer ID does not exist, the API should return http status code 404 (Not Found)
- [x] In case of an unexpected error, the API should return a status code 500 (Internal Server Error) along with the error message
