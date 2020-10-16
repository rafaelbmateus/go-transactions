package transaction

import "time"

func NewFixtureTransaction() *Transaction {
	return &Transaction{
		ID:              1,
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          123.45,
		EventDate:       time.Now(),
	}
}
