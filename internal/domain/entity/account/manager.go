package account

type manager struct {
	repo repository
}

// NewManager create new manager
func NewManager(r repository) *manager {
	return &manager{
		repo: r,
	}
}

// Create an account
func (s *manager) Create(e *Account) (*Account, error) {
	return s.repo.Create(e)
}

// Get an account
func (s *manager) Get(id int) (*Account, error) {
	return s.repo.Get(id)
}

// List accounts
func (s *manager) List() ([]*Account, error) {
	return s.repo.List()
}

// Update a account
func (s *manager) Update(e *Account) error {
	return s.repo.Update(e)
}

// Delete a account
func (s *manager) Delete(id int) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
