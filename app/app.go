package app

import (
	"net/http"

	"github.com/codacube/go-microservice-banking/domain"
	"github.com/codacube/go-microservice-banking/service"
	"github.com/gorilla/mux"
)

// Exported functions must start with a capital letter, otherwise they will not be accessible from other packages!
func Start() {

	router := mux.NewRouter()

	// Wiring it up
	// ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// Define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// Start server
	http.ListenAndServe("localhost:8000", router)
}
