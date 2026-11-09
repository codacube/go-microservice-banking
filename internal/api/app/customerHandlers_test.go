package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/codacube/banking-lib/errs"
	"github.com/codacube/go-microservice-banking/internal/api/domain"
	"github.com/codacube/go-microservice-banking/internal/api/dto"
	"github.com/codacube/go-microservice-banking/internal/api/mocks/service"
	"github.com/gorilla/mux"
	"go.uber.org/mock/gomock"
)

func setup(t *testing.T) (*mux.Router, *service.MockCustomerService) {
	t.Helper()
	ctrl := gomock.NewController(t)
	mockService := service.NewMockCustomerService(ctrl)
	ch := CustomerHandlers{service: mockService}
	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer)
	return router, mockService
}

func Test_should_return_customers_with_status_code_200(t *testing.T) {
	// Arrange
	router, mockService := setup(t)

	dummyCustomers := []dto.CustomerResponse{
		{"1001", "Ashish", "New Delhi", "110011", "2000-01-01", "1"},
		{"1002", "Rob", "New Delhi", "110011", "2000-01-01", "1"},
	}
	mockService.EXPECT().GetAllCustomers(domain.CustomerStatus("")).Return(dummyCustomers, nil)
	request := httptest.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Errorf("status code = %d, want %d", recorder.Code, http.StatusOK)
	}
	if got := recorder.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want %q", got, "application/json")
	}
	var gotCustomers []dto.CustomerResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &gotCustomers); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	if !reflect.DeepEqual(gotCustomers, dummyCustomers) {
		t.Errorf("response body = %#v, want %#v", gotCustomers, dummyCustomers)
	}
}

func Test_should_filter_customers_by_status(t *testing.T) {
	tests := []struct {
		name   string
		status domain.CustomerStatus
	}{
		{name: "active", status: domain.CustomerStatusActive},
		{name: "inactive", status: domain.CustomerStatusInactive},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router, mockService := setup(t)
			mockService.EXPECT().GetAllCustomers(tt.status).Return([]dto.CustomerResponse{}, nil)
			request := httptest.NewRequest(http.MethodGet, "/customers?status="+string(tt.status), nil)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, request)

			if recorder.Code != http.StatusOK {
				t.Errorf("status code = %d, want %d", recorder.Code, http.StatusOK)
			}
			if got := recorder.Header().Get("Content-Type"); got != "application/json" {
				t.Errorf("Content-Type = %q, want %q", got, "application/json")
			}
			var gotCustomers []dto.CustomerResponse
			if err := json.Unmarshal(recorder.Body.Bytes(), &gotCustomers); err != nil {
				t.Fatalf("failed to decode response body: %v", err)
			}
			if len(gotCustomers) != 0 {
				t.Errorf("response body = %#v, want an empty customer list", gotCustomers)
			}
		})
	}
}

func Test_should_return_status_code_400_for_invalid_customer_status(t *testing.T) {
	router, _ := setup(t)
	request := httptest.NewRequest(http.MethodGet, "/customers?status=unknown", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("status code = %d, want %d", recorder.Code, http.StatusBadRequest)
	}
	if got := recorder.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want %q", got, "application/json")
	}
	var gotError struct {
		Message string `json:"message"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &gotError); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	if want := "Status must be active or inactive"; gotError.Message != want {
		t.Errorf("error message = %q, want %q", gotError.Message, want)
	}
}

func Test_should_return_status_code_500_with_error_message(t *testing.T) {
	// Arrange
	router, mockService := setup(t)
	mockService.EXPECT().GetAllCustomers(domain.CustomerStatus("")).Return(nil, errs.NewUnexpectedError("some database error"))
	request := httptest.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("status code = %d, want %d", recorder.Code, http.StatusInternalServerError)
	}
	if got := recorder.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want %q", got, "application/json")
	}
	var gotError struct {
		Message string `json:"message"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &gotError); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	if want := "some database error"; gotError.Message != want {
		t.Errorf("error message = %q, want %q", gotError.Message, want)
	}
}

func Test_should_return_customer_with_status_code_200(t *testing.T) {
	// Arrange
	router, mockService := setup(t)

	dummyCustomer := dto.CustomerResponse{"1001", "Ashish", "New Delhi", "110011", "2000-01-01", "1"}
	mockService.EXPECT().GetCustomer("1001").Return(&dummyCustomer, nil)
	request := httptest.NewRequest(http.MethodGet, "/customers/1001", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Errorf("status code = %d, want %d", recorder.Code, http.StatusOK)
	}
	if got := recorder.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want %q", got, "application/json")
	}
	var gotCustomer dto.CustomerResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &gotCustomer); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	if !reflect.DeepEqual(gotCustomer, dummyCustomer) {
		t.Errorf("response body = %#v, want %#v", gotCustomer, dummyCustomer)
	}
}

func Test_should_return_status_code_500_with_error_message_for_getCustomer(t *testing.T) {
	// Arrange
	router, mockService := setup(t)
	mockService.EXPECT().GetCustomer("1001").Return(nil, errs.NewUnexpectedError("some database error"))
	request := httptest.NewRequest(http.MethodGet, "/customers/1001", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("status code = %d, want %d", recorder.Code, http.StatusInternalServerError)
	}
	if got := recorder.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want %q", got, "application/json")
	}
	var gotError struct {
		Message string `json:"message"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &gotError); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	if want := "some database error"; gotError.Message != want {
		t.Errorf("error message = %q, want %q", gotError.Message, want)
	}
}
