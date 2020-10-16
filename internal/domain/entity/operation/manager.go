package operation

type manager struct {
	repo repository
}

// NewManager create new manager
func NewManager(r repository) *manager {
	return &manager{
		repo: r,
	}
}

// Create an operation type
func (s *manager) Create(e *OperationType) error {
	// e.CreatedAt = time.Now()
	return s.repo.Create(e)
}

// Get an operation type
func (s *manager) Get(id int) (*OperationType, error) {
	return s.repo.Get(id)
}

// List the operation types
func (s *manager) List() ([]*OperationType, error) {
	return s.repo.List()
}

// Delete an operation type
func (s *manager) Delete(id int) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

// Update an operation type
func (s *manager) Update(e *OperationType) error {
	return s.repo.Update(e)
}
