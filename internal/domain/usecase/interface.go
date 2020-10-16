package usecase

import (
	"github.com/rafaelbmateus/go-transactions/internal/domain/entity/account"
	"github.com/rafaelbmateus/go-transactions/internal/domain/entity/transaction"
)

// UseCase use case interface
type UseCase interface {
	NewAccount(*account.Account) (*account.Account, error)
	GetAccount(id int) (*account.Account, error)
	RegisterTransaction(*transaction.Transaction) error
}
