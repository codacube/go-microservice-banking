module github.com/codacube/go-microservice-banking

go 1.27.0

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.10.1
	github.com/golang-jwt/jwt/v4 v4.5.2
	github.com/gorilla/mux v1.8.1
	github.com/jmoiron/sqlx v1.4.0
	github.com/joho/godotenv v1.5.1
	go.uber.org/mock v0.6.0
	go.uber.org/zap v1.28.0
)

require (
	filippo.io/edwards25519 v1.2.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/mod v0.27.0 // indirect
	golang.org/x/sync v0.16.0 // indirect
	golang.org/x/tools v0.36.0 // indirect
)

tool go.uber.org/mock/mockgen
