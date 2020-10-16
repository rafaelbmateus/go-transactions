package operation

import (
	"errors"
)

// iRepo in memory repo
type iRepo struct {
	m map[int]*OperationType
}

// NewInmemRepository create new repository
func NewInmemRepository() *iRepo {
	var m = map[int]*OperationType{}
	return &iRepo{
		m: m,
	}
}

// Create an operation type
func (r *iRepo) Create(e *OperationType) error {
	r.m[e.ID] = e
	return nil
}

// Get an operation type
func (r *iRepo) Get(id int) (*OperationType, error) {
	if r.m[id] == nil {
		return nil, errors.New("Not found")
	}
	return r.m[id], nil
}

// Update an operation type
func (r *iRepo) Update(e *OperationType) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

// List operation types
func (r *iRepo) List() ([]*OperationType, error) {
	var d []*OperationType
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete an operation type
func (r *iRepo) Delete(id int) error {
	if r.m[id] == nil {
		return errors.New("Not found")
	}
	r.m[id] = nil
	return nil
}
