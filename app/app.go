package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Exported functions must start with a capital letter, otherwise they will not be accessible from other packages!
func Start() {

	mux := mux.NewRouter()

	// Define routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	// Start server
	http.ListenAndServe("localhost:8000", mux)
}
