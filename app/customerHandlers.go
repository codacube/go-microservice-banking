package app

import (
	"encoding/json"
	"net/http"

	"github.com/codacube/go-microservice-banking/domain"
	"github.com/codacube/go-microservice-banking/errs"
	"github.com/codacube/go-microservice-banking/service"
	"github.com/gorilla/mux"
)

// User/Driving/Primary Side - Adapter
// CustomerHandlers is the inbound HTTP adapter for customer use cases. It
// depends on the primary CustomerService port, not on data-storage details.
type CustomerHandlers struct {
	service service.CustomerService
}

// HTTP handler function for the /customers endpoint. It uses the CustomerService
// primary port to retrieve customer data and respond to the client.
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	var status domain.CustomerStatus
	switch r.URL.Query().Get("status") {
	case "":
	case string(domain.CustomerStatusActive):
		status = domain.CustomerStatusActive
	case string(domain.CustomerStatusInactive):
		status = domain.CustomerStatusInactive
	default:
		writeResponse(w, http.StatusBadRequest, errs.NewBadRequestError("Status must be active or inactive").AsMessage())
		return
	}

	customers, err := ch.service.GetAllCustomers(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	// Extract the customer ID from the URL path
	vars := mux.Vars(r)
	customerID := vars["customer_id"]

	// Pass the ID to the service layer to retrieve the customer
	customer, err := ch.service.GetCustomer(customerID)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
