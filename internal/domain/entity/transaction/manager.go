package transaction

import (
	"time"
)

type manager struct {
	repo repository
}

// NewManager create new manager
func NewManager(r repository) *manager {
	return &manager{
		repo: r,
	}
}

// Create a transaction
func (s *manager) Create(e *Transaction) error {
	e.EventDate = time.Now()
	return s.repo.Create(e)
}

// Get a account
func (s *manager) Get(id int) (*Transaction, error) {
	return s.repo.Get(id)
}

// List accounts
func (s *manager) List() ([]*Transaction, error) {
	return s.repo.List()
}
