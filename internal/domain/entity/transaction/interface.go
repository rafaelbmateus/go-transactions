package transaction

// Reader transaction reader interface
type Reader interface {
	Get(id int) (*Transaction, error)
	List() ([]*Transaction, error)
}

// Writer transaction writer interface
type Writer interface {
	Create(e *Transaction) error
}

// Manager interface
type Manager interface {
	repository
}

// repository interface
type repository interface {
	Reader
	Writer
}
