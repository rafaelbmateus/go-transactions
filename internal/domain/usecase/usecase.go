package usecase

import (
	"errors"

	"github.com/rafaelbmateus/go-transactions/internal/domain/entity/account"
	"github.com/rafaelbmateus/go-transactions/internal/domain/entity/operation"
	"github.com/rafaelbmateus/go-transactions/internal/domain/entity/transaction"
)

type usecase struct {
	accountManager     account.Manager
	operationManager   operation.Manager
	transactionManager transaction.Manager
}

// NewUseCase create new use case
func NewUseCase(a account.Manager, o operation.Manager, t transaction.Manager) *usecase {
	return &usecase{
		accountManager:     a,
		operationManager:   o,
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
	account, err := s.accountManager.Get(t.AccountID)
	if err != nil {
		return err
	}

	operation, err := s.operationManager.Get(t.OperationTypeID)
	if err != nil {
		return err
	}

	if operation.IsDebit {
		if account.AvailableCreditLimit <= t.Amount {
			return errors.New("No credit limit available")
		}
		account.AvailableCreditLimit -= t.Amount
		err = s.accountManager.Update(account)
		if err != nil {
			return err
		}
	}
	err = s.transactionManager.Create(t)
	if err != nil {
		return err
	}
	return nil
}

// GetTransactions get all transactions
func (s *usecase) GetTransactions() ([]*transaction.Transaction, error) {
	transactions, err := s.transactionManager.List()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
