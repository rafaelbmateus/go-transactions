package operation

// OperationType represets one type of operation used in the trasactions
type OperationType struct {
	ID          int
	Description string
	IsDebit     bool
}
