package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/anucha-tk/go_bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	Currency      string `json:"currency" binding:"required,currency"`
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountId   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !server.validAccount(ctx, req.FromAccountID, req.Currency) {
		return
	}
	if !server.validAccount(ctx, req.ToAccountId, req.Currency) {
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountId,
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountId int64, currency string) bool {
	account, err := server.store.GetAccount(ctx, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency missmatch: %s vs %s", account.ID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}
	return true
}
