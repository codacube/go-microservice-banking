# To Build / Run

```go
go build -o bin/api ./cmd/api

go run ./cmd/api
```

## Summary of the internal/api packages

- app: this is the HTTP/bootstrap layer. It creates the router, loads the DB client, wires repositories into handlers, and exposes the API endpoints. See `app.go`.
  - It starts the service and registers routes for customers and accounts.

- domain: this is the core business model layer. It contains the main domain objects and repository contracts for customer and account operations. See:
  - `customer.go`
  - `account.go`
  - This package defines models like Customer and Account, plus business rules such as status and withdrawal validation.

- dto: this is the data-transfer layer for request/response payloads. It defines the JSON shapes sent to and returned from the API. See:
  - `customerResponse.go`
  - There are also account-related DTOs in the same folder for creating accounts and transactions.

- service: this is the business logic layer. It orchestrates the use cases exposed by the API, such as retrieving customers and creating accounts/transactions. See:
  - `customerService.go`
  - `accountService.go`
  - These services depend on repository interfaces from the domain layer and convert domain objects into DTOs for the API.

### Overall flow

app → service → domain
with dto as the contract for API payloads.

This means the HTTP layer does not talk directly to the database; it delegates to service logic, which in turn uses repository abstractions defined in the domain layer.
