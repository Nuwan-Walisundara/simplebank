package api

import (
	"fmt"
	"net/http"

	db "github.com/Nuwan-Walisundara/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateTransferReq struct {
	FromAccID int64  `json:"fromAccId" binding:"required"`
	ToAccId   int64  `json:"toAccId" binding:"required"`
	Amount    int64  `json:"amount" binding:"required,gt=0"`
	Currency  string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var createTransferReq CreateTransferReq

	if err := ctx.ShouldBindJSON(&createTransferReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	val := server.validateAccount(ctx, createTransferReq.FromAccID, createTransferReq.ToAccId)

	if !val {
		return
	}

	result, err := server.store.TransferTx(ctx, db.TransferParams{FromAccountId: createTransferReq.FromAccID,
		ToAccountID: createTransferReq.ToAccId,
		Balance:     createTransferReq.Amount,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)

}

func (server *Server) validateAccount(ctx *gin.Context, fromAccountId int64, toAccountId int64) bool {
	var accountIds []int32

	accountIds = append(accountIds, int32(fromAccountId), int32(toAccountId))

	accounts, err := server.store.GetAccountsByIDs(ctx, accountIds)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if len(accounts) > 1 && accounts[0].Currency == accounts[1].Currency {
		return true
	} else if len(accounts) > 1 && accounts[0].Currency != accounts[1].Currency {
		err := fmt.Errorf("Account currency doesnot match acc1 :%s acc2: %s", accounts[0].Currency, accounts[1].Currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	} else {
		err := fmt.Errorf("either one of accounnt is doesnot exists ")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}
	return false
}
