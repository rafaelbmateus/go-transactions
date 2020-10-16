package presenter

import "time"

// Transaction represents a financial transaction
type Transaction struct {
	ID              int       `json:"id"`
	AccountID       int       `json:"account_id"`
	OperationTypeID int       `json:"operation_type_id"`
	Amount          float64   `json:"amount"`
	EventDate       time.Time `json:"event_date"`
}
