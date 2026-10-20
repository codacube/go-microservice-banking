package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codacube/go-microservice-banking/domain"
	"github.com/codacube/go-microservice-banking/service"
)

func TestGetAllCustomersStatusFilter(t *testing.T) {
	handler := CustomerHandlers{
		service: service.NewCustomerService(domain.NewCustomerRepositoryStub()),
	}

	request := httptest.NewRequest(http.MethodGet, "/customers?status=active", nil)
	response := httptest.NewRecorder()
	handler.getAllCustomers(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status code = %d, want %d", response.Code, http.StatusOK)
	}

	var customers []struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(response.Body).Decode(&customers); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(customers) != 2 {
		t.Fatalf("got %d customers, want 2", len(customers))
	}
	for _, customer := range customers {
		if customer.Status != string(domain.CustomerStatusActive) {
			t.Errorf("customer status = %q, want %q", customer.Status, domain.CustomerStatusActive)
		}
	}
}

func TestGetAllCustomersRejectsUnsupportedStatus(t *testing.T) {
	handler := CustomerHandlers{
		service: service.NewCustomerService(domain.NewCustomerRepositoryStub()),
	}

	request := httptest.NewRequest(http.MethodGet, "/customers?status=unknown", nil)
	response := httptest.NewRecorder()
	handler.getAllCustomers(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("status code = %d, want %d", response.Code, http.StatusBadRequest)
	}
}
