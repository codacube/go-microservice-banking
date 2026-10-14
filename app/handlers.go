package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name     string `json:"full_name" xml:"name"`
	City     string `json:"city" xml:"city"`
	Postcode string `json:"post_code" xml:"post_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "John", City: "Auckland", Postcode: "10001"},
		{Name: "Jane", City: "Wellington", Postcode: "90001"},
		{Name: "Bob", City: "Christchurch", Postcode: "60601"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
