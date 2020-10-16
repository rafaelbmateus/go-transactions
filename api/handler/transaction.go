package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelbmateus/go-transactions/api/presenter"
	"github.com/rafaelbmateus/go-transactions/internal/domain/entity/transaction"
	"github.com/rafaelbmateus/go-transactions/internal/domain/usecase"
)

// TransactionRouters is the endpoints of the server
func TransactionRouters(e *gin.Engine, u usecase.UseCase) {
	e.POST("/transactions", RegisterTransaction(u))
}

// RegisterTransaction register a new transaction
func RegisterTransaction(u usecase.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var t presenter.Transaction
		err := c.BindJSON(&t)
		if err != nil {
			responseFailure(c, http.StatusText(http.StatusInternalServerError),
				"account can't be created",
				"error when converting the parameters sent to json", "", http.StatusInternalServerError)
			return
		}

		err = u.RegisterTransaction(&transaction.Transaction{
			ID:              t.ID,
			AccountID:       t.AccountID,
			OperationTypeID: t.OperationTypeID,
			Amount:          t.Amount,
		})
		if err != nil {
			responseFailure(c, http.StatusText(http.StatusInternalServerError),
				"account can't be created",
				"internal server error when creating a new account", "", http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
	}
}
