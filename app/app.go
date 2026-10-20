package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codacube/go-microservice-banking/domain"
	"github.com/codacube/go-microservice-banking/service"
	"github.com/gorilla/mux"
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
	// ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// Define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// Start server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
