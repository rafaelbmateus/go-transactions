package usecase

import (
	"testing"

	"github.com/rafaelbmateus/go-transactions/internal/domain/entity/transaction"

	"github.com/go-playground/assert"
	"github.com/golang/mock/gomock"
	"github.com/rafaelbmateus/go-transactions/internal/domain/entity/account"
	amock "github.com/rafaelbmateus/go-transactions/internal/domain/entity/account/mock"
	tmock "github.com/rafaelbmateus/go-transactions/internal/domain/entity/transaction/mock"
)

func TestNewAccount(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	aMock := amock.NewMockManager(controller)
	tMock := tmock.NewMockManager(controller)
	uc := NewUseCase(aMock, tMock)
	t.Run("with success", func(t *testing.T) {
		a := account.NewFixtureAccount()
		aMock.EXPECT().Create(a).Return(&account.Account{ID: a.ID, DocumentNumber: a.DocumentNumber}, nil)
		acc, err := uc.NewAccount(a)
		assert.Equal(t, nil, err)
		assert.Equal(t, acc.ID, a.ID)
		assert.Equal(t, acc.DocumentNumber, a.DocumentNumber)
	})
}

func TestGetAccount(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	aMock := amock.NewMockManager(controller)
	tMock := tmock.NewMockManager(controller)
	uc := NewUseCase(aMock, tMock)
	t.Run("with success", func(t *testing.T) {
		a := account.NewFixtureAccount()
		aMock.EXPECT().Get(a.ID).Return(&account.Account{ID: a.ID, DocumentNumber: a.DocumentNumber}, nil)
		acc, err := uc.GetAccount(a.ID)
		assert.Equal(t, nil, err)
		assert.Equal(t, acc.ID, a.ID)
		assert.Equal(t, acc.DocumentNumber, a.DocumentNumber)
	})
}

func TestRegisterTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	aMock := amock.NewMockManager(controller)
	tMock := tmock.NewMockManager(controller)
	uc := NewUseCase(aMock, tMock)
	t.Run("with success", func(t *testing.T) {
		t1 := transaction.NewFixtureTransaction()
		tMock.EXPECT().Create(t1).Return(nil)
		err := uc.RegisterTransaction(t1)
		assert.Equal(t, nil, err)
		assert.Equal(t, t1.Amount, 123.45)
	})
}

func TestGetTransactions(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	aMock := amock.NewMockManager(controller)
	tMock := tmock.NewMockManager(controller)
	uc := NewUseCase(aMock, tMock)
	t.Run("with success", func(t *testing.T) {
		t1 := transaction.NewFixtureTransaction()
		t2 := transaction.NewFixtureTransaction()
		tMock.EXPECT().List().Return(
			[]*transaction.Transaction{
				&transaction.Transaction{
					ID:              t1.ID,
					AccountID:       t1.AccountID,
					OperationTypeID: t1.OperationTypeID,
					Amount:          t1.Amount,
					EventDate:       t1.EventDate,
				},
				&transaction.Transaction{
					ID:              t2.ID,
					AccountID:       t2.AccountID,
					OperationTypeID: t2.OperationTypeID,
					Amount:          t2.Amount,
					EventDate:       t2.EventDate,
				},
			}, nil)
		transactions, err := uc.GetTransactions()
		assert.Equal(t, nil, err)
		assert.Equal(t, len(transactions), 2)
	})
}
