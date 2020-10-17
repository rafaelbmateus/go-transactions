package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafaelbmateus/go-transactions/api/presenter"
	"github.com/rafaelbmateus/go-transactions/internal/domain/entity/account"
	"github.com/rafaelbmateus/go-transactions/internal/domain/usecase"
)

// AccountRouters is the endpoints of the server
func AccountRouters(e *gin.Engine, u usecase.UseCase) {
	e.POST("/accounts", CreateAccount(u))
	e.GET("/accounts/:accountId", GetAccount(u))
}

// CreateAccount create a new account
func CreateAccount(u usecase.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var a presenter.Account
		err := c.BindJSON(&a)
		if err != nil {
			responseFailure(c, http.StatusText(http.StatusInternalServerError),
				"Account can't be created",
				"Error when converting the parameters sent to json", "", http.StatusInternalServerError)
			return
		}

		acc, err := u.NewAccount(&account.Account{
			ID: a.ID, DocumentNumber: a.DocumentNumber,
			AvailableCreditLimit: a.AvailableCreditLimit,
		})
		if err != nil {
			responseFailure(c, http.StatusText(http.StatusInternalServerError),
				"Account can't be created",
				fmt.Sprintf("Internal server error when creating a new account - datails err: %s", err.Error()), "", http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusCreated, acc)
	}
}

// GetAccount get an existing account
func GetAccount(u usecase.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		accountID, err := strconv.Atoi(c.Param("accountId"))
		if err != nil {
			responseFailure(c, http.StatusText(http.StatusInternalServerError),
				"Can't get an account",
				"The parameter accountId it's not an interger type identifier", "", http.StatusBadRequest)
			return
		}

		account, err := u.GetAccount(accountID)
		if err != nil {
			responseFailure(c, http.StatusText(http.StatusInternalServerError),
				"Can't get an account - There was an internal error when obtaining an account",
				err.Error(), "", http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, account)
	}
}
