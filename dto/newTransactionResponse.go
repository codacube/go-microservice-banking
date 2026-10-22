package dto

type NewTransactionResponse struct {
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
}
