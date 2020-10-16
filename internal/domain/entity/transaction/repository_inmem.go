package transaction

import (
	"errors"
)

// iRepo in memory repo
type iRepo struct {
	m map[int]*Transaction
}

// NewInmemRepository create new repository
func NewInmemRepository() *iRepo {
	var m = map[int]*Transaction{}
	return &iRepo{
		m: m,
	}
}

// Create a transaction
func (r *iRepo) Create(e *Transaction) error {
	r.m[e.ID] = e
	return nil
}

// Get a transaction
func (r *iRepo) Get(id int) (*Transaction, error) {
	if r.m[id] == nil {
		return nil, errors.New("Not found")
	}
	return r.m[id], nil
}

// List transactions
func (r *iRepo) List() ([]*Transaction, error) {
	var d []*Transaction
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}
