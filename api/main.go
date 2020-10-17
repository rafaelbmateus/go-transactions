package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelbmateus/go-transactions/api/handler"
	"github.com/rafaelbmateus/go-transactions/internal/domain/entity/account"
	"github.com/rafaelbmateus/go-transactions/internal/domain/entity/operation"
	"github.com/rafaelbmateus/go-transactions/internal/domain/entity/transaction"
	"github.com/rafaelbmateus/go-transactions/internal/domain/usecase"
)

func main() {
	accountRepo := account.NewInmemRepository()
	accountManager := account.NewManager(accountRepo)

	operationRepo := operation.NewInmemRepository()
	operationManager := operation.NewManager(operationRepo)
	defaultOperations(operationManager)

	transactionRepo := transaction.NewInmemRepository()
	transactionManager := transaction.NewManager(transactionRepo)

	usecase := usecase.NewUseCase(accountManager, operationManager, transactionManager)

	r := gin.Default()

	// healthy
	handler.HealthyRouters(r)

	// account
	handler.AccountRouters(r, usecase)

	// transaction
	handler.TransactionRouters(r, usecase)

	err := r.Run(":3000")
	if err != nil {
		panic("error to run the web service")
	}
}

func defaultOperations(o operation.Manager) {
	_ = o.Create(&operation.OperationType{ID: 1, Description: "Compra a vista", IsDebit: true})
	_ = o.Create(&operation.OperationType{ID: 2, Description: "Compra parcelada", IsDebit: true})
	_ = o.Create(&operation.OperationType{ID: 3, Description: "Saque", IsDebit: true})
	_ = o.Create(&operation.OperationType{ID: 4, Description: "Pagamento", IsDebit: false})
}
