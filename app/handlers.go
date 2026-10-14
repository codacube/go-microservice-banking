package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/codacube/go-microservice-banking/service"
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
