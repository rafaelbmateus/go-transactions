package account

// Reader interface
type Reader interface {
	Get(id int) (*Account, error)
	List() ([]*Account, error)
}

// Writer Account writer
type Writer interface {
	Create(e *Account) (*Account, error)
	Update(e *Account) error
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
