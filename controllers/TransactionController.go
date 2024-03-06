package controllers

import (
	db "bankassignment/database"
	"bankassignment/models"
	"bankassignment/queries"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(ctx *gin.Context) {
	tx, txErr := db.DB.Begin()
	if txErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": txErr.Error()})
		tx.Rollback()
		return
	}
	var transaction *models.Transaction
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}
	switch transaction.TransactionType {
	case "Debit":
		_, err := queries.Debit(float32(transaction.AccountId), float32(transaction.TransactionAmount), tx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			tx.Rollback()
			return
		}
	case "Credit":
		_, err := queries.Credit(float32(transaction.AccountId), float32(transaction.TransactionAmount), tx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			tx.Rollback()
			return
		}
	case "Transfer":
		_, err := queries.Debit(float32(transaction.AccountId), float32(transaction.TransactionAmount), tx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			tx.Rollback()
			return
		}
		_, err2 := queries.Credit(float32(transaction.RecieverAccountId), float32(transaction.TransactionAmount), tx)
		if err2 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
			tx.Rollback()
			return
		}
	}

	rows_affected, err := queries.CreateTransaction(transaction, tx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{"Transaction inserted successfully.Number of insertions": rows_affected})
}

func GetTransactionByAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	transactions, err := queries.GetTransactionByAccount(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Transactions are as follows": transactions})
}

func GetTransactionDetailsById(ctx *gin.Context) {
	id := ctx.Param("id")
	transaction, err := queries.GetTransactionDetailsById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Transaction :": transaction})
}
