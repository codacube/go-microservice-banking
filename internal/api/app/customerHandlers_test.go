package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codacube/go-microservice-banking/internal/api/domain"
	"github.com/codacube/go-microservice-banking/internal/api/dto"
	"github.com/codacube/go-microservice-banking/internal/api/service"
	"github.com/codacube/go-microservice-banking/internal/errs"
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
	return router, mockService
}

func Test_should_return_customers_with_status_code_200(t *testing.T) {
	router, mockService := setup(t)

	dummyCustomers := []dto.CustomerResponse{
		{"1001", "Ashish", "New Delhi", "110011", "2000-01-01", "1"},
		{"1002", "Rob", "New Delhi", "110011", "2000-01-01", "1"},
	}
	mockService.EXPECT().GetAllCustomers(domain.CustomerStatus("")).Return(dummyCustomers, nil)
	request := httptest.NewRequest(http.MethodGet, "/customers", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Errorf("status code = %d, want %d", recorder.Code, http.StatusOK)
	}
}

func Test_should_return_status_code_500_with_error_message(t *testing.T) {
	router, mockService := setup(t)
	mockService.EXPECT().GetAllCustomers(domain.CustomerStatus("")).Return(nil, errs.NewUnexpectedError("some database error"))
	request := httptest.NewRequest(http.MethodGet, "/customers", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("status code = %d, want %d", recorder.Code, http.StatusInternalServerError)
	}
}
