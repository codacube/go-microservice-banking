package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/codacube/go-microservice-banking/service"
	"github.com/gorilla/mux"
)

type Customer struct {
	Name     string `json:"full_name" xml:"name"`
	City     string `json:"city" xml:"city"`
	Postcode string `json:"post_code" xml:"post_code"`
}

// User/Driving/Primary Side - Adapter
// CustomerHandlers is the inbound HTTP adapter for customer use cases. It
// depends on the primary CustomerService port, not on data-storage details.
type CustomerHandlers struct {
	service service.CustomerService
}

// HTTP handler function for the /customers endpoint. It uses the CustomerService
// primary port to retrieve customer data and respond to the client.
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "John", City: "Auckland", Postcode: "10001"},
	// 	{Name: "Jane", City: "Wellington", Postcode: "90001"},
	// 	{Name: "Bob", City: "Christchurch", Postcode: "60601"},
	// }

	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		http.Error(w, "Failed to retrieve customers", http.StatusInternalServerError)
		return
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	// Extract the customer ID from the URL path
	vars := mux.Vars(r)
	customerID := vars["customer_id"]

	// Pass the ID to the service layer to retrieve the customer
	customer, err := ch.service.GetCustomer(customerID)
	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprint(w, err.Message)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
