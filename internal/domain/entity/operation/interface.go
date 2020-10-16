package operation

// Reader interface
type Reader interface {
	Get(id int) (*OperationType, error)
	List() ([]*OperationType, error)
}

// Writer Account writer
type Writer interface {
	Create(e *OperationType) error
	Update(e *OperationType) error
	Delete(id int) error
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
