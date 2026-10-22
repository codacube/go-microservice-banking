package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codacube/go-microservice-banking/domain"
	"github.com/codacube/go-microservice-banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	// Check if critical variables exist
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variables not defined: SERVER_ADDRESS or SERVER_PORT is missing")
	}
	if os.Getenv("DB_ADDR") == "" || os.Getenv("DB_PORT") == "" || os.Getenv("DB_NAME") == "" {
		log.Fatal("Environment variables not defined: Database server details missing")
	}
	if os.Getenv("DB_USER") == "" || os.Getenv("DB_PASS") == "" {
		log.Fatal("Environment variables not defined: Database credentials missing")
	}
}

// Exported functions must start with a capital letter, otherwise they will not be accessible from other packages!
func Start() {
	sanityCheck()

	router := mux.NewRouter()

	// Wiring it up
	dbClient := getDBClient()

	// ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandlers{service: service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service: service.NewAccountService(accountRepositoryDb)}

	// Define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.NewTransaction).Methods(http.MethodPost)

	// Start server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func getDBClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASS")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// The appended flags might be needed to for localhost development, but appears to work without it
	// sql.Open("mysql", "root:your_strong_password@tcp(localhost:3306)/banking?allowPublicKeyRetrieval=true&parseTime=true&columnsWithAlias=true")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	// dataSource := fmt.Sprintf("%s:%s@tcp(localhost:3306)/banking", dbUser, dbPasswd)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
