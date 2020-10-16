package account

import (
	"errors"
)

// iRepo in memory repo
type iRepo struct {
	m map[int]*Account
}

// NewInmemRepository create new repository
func NewInmemRepository() *iRepo {
	var m = map[int]*Account{}
	return &iRepo{
		m: m,
	}
}

// Create an account
func (r *iRepo) Create(a *Account) (*Account, error) {
	a.ID = len(r.m) + 1
	r.m[a.ID] = a
	return a, nil
}

// Get an account
func (r *iRepo) Get(id int) (*Account, error) {
	if r.m[id] == nil {
		return nil, errors.New("Not found")
	}
	return r.m[id], nil
}

// Update an account
func (r *iRepo) Update(e *Account) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

// List accounts
func (r *iRepo) List() ([]*Account, error) {
	var d []*Account
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete an account
func (r *iRepo) Delete(id int) error {
	if r.m[id] == nil {
		return errors.New("Not found")
	}
	r.m[id] = nil
	return nil
}
