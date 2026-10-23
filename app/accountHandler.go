package app

import (
	"encoding/json"
	"net/http"

	"github.com/codacube/go-microservice-banking/dto"
	"github.com/codacube/go-microservice-banking/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customer_id"]

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerID = customerID
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

func (h AccountHandler) NewTransaction(w http.ResponseWriter, r *http.Request) {
	// Get the account_id and customer_id from the URL
	vars := mux.Vars(r)
	accountID := vars["account_id"]
	customerID := vars["customer_id"]

	// Decode incoming request
	var request dto.NewTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {

		// Build the request object
		request.AccountID = accountID
		request.CustomerID = customerID

		// Save the new transaction
		transaction, appError := h.service.NewTransaction(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, transaction)
		}
	}

}
