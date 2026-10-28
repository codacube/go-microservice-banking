TODO

- Add git
- Finish files
- use .env file for environment variables
- Hash the user passwords

# To Build / Run

```go
go build -o bin/api ./cmd/api

go run ./cmd/api
```

# ## Summary of the internal/auth packages

- app: this is the startup/bootstrap layer. It sets up the HTTP router, loads environment config, creates the DB client, and registers the auth endpoints. See `app.go`.
  - It wires the service into routes like login, verify, and a placeholder register endpoint.

- domain: this is the core auth/business layer. It contains the repository contract, login model, JWT claims/token logic, and role/permission rules that determine which routes a user can access. The main implementation is in `authRepository.go`, with supporting auth-related types under the same package.

- dto: this is the data-transfer layer. It defines the JSON request/response shapes used between clients and the service, such as login payloads and token-related payloads. The request model is defined in `loginRequest.go`.

- service: this is the use-case/business logic layer. It implements auth operations like:
  - login
  - token verification
  - refresh-token flow
  - permission checks against the role system
  - This logic lives in `authService.go`.

In short, the package flow is:

app -> service -> domain
and dto is the contract layer between HTTP and service code.
