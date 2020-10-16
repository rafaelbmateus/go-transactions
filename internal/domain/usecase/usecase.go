package usecase

import (
	"github.com/rafaelbmateus/go-transactions/internal/domain/entity/account"
	"github.com/rafaelbmateus/go-transactions/internal/domain/entity/transaction"
)

type usecase struct {
	accountManager     account.Manager
	transactionManager transaction.Manager
}

// NewUseCase create new use case
func NewUseCase(a account.Manager, t transaction.Manager) *usecase {
	return &usecase{
		accountManager:     a,
		transactionManager: t,
	}
}

// Create an new account
func (s *usecase) NewAccount(a *account.Account) (*account.Account, error) {
	acc, err := s.accountManager.Create(a)
	// validate the fields and business rule

	if err != nil {
		return nil, err
	}
	return acc, nil
}

// Get an existing account
func (s *usecase) GetAccount(id int) (*account.Account, error) {
	a, err := s.accountManager.Get(id)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// RegisterTransaction register a new transaction
func (s *usecase) RegisterTransaction(t *transaction.Transaction) error {
	err := s.transactionManager.Create(t)
	if err != nil {
		return err
	}
	return nil
}
