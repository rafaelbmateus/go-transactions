package transaction

import (
	"time"
)

// Transaction represents a financial transaction
type Transaction struct {
	ID              int
	AccountID       int
	OperationTypeID int
	Amount          float64
	EventDate       time.Time
}
