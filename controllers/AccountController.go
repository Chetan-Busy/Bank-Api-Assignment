package controllers

import (
	db "bankassignment/database"
	"bankassignment/models"
	"bankassignment/queries"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAccount(ctx *gin.Context) {
	var account *models.Account
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rows_affected, err := queries.CreateAccount(account)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"Number of rows affected": rows_affected})
}

func DeleteAccountById(ctx *gin.Context) {
	id := ctx.Param("id")
	no_of_rows_deleted, err := queries.DeleteAccountById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"No of rows deleted:": no_of_rows_deleted})
}

func GetAccountDetailsById(ctx *gin.Context) {
	id := ctx.Param("id")
	account, err := queries.GetAccountDetailsById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"The account details are": account})
}

func UpdateAccount(ctx *gin.Context) {
	var account *models.Account
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rows_updated, err := queries.UpdateAccount(account)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if rows_updated == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No account found with that id"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Account updated successfully.Number of rows updated": rows_updated})
}

// {
// accountId -> float32
// amount -> float32
// }
func Credit(ctx *gin.Context) {
	tx, txErr := db.DB.Begin()
	if txErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": txErr.Error()})
		tx.Rollback()
		return
	}
	var jsonData map[string]float32
	if err := ctx.ShouldBindJSON(&jsonData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	accountId := float32(jsonData["accountId"])
	amount := float32(jsonData["amount"])

	rows_updated, err := queries.Credit(accountId, amount, tx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}
	if rows_updated == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No account with the given account Id found"})
		tx.Rollback()
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{"Amount credit successfull. Number of account updated : ": rows_updated})
}

func Debit(ctx *gin.Context) {
	tx, txErr := db.DB.Begin()
	if txErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": txErr.Error()})
		tx.Rollback()
		return
	}
	var jsonData map[string]float32
	if err := ctx.ShouldBindJSON(&jsonData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	accountId := float32(jsonData["accountId"])
	amount := float32(jsonData["amount"])
	rows_updated, err := queries.Debit(accountId, amount, tx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}
	if rows_updated == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No account with the given account Id found"})
		tx.Rollback()
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{"Amount debited successfull. Number of account updated : ": rows_updated})

}

//	{
//		senderAccountId -> float32
//		RecieverAccountId -> float32
//		Amount -> float32
//	}
func Transfer(ctx *gin.Context) {
	tx, txErr := db.DB.Begin()
	if txErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": txErr.Error()})
		tx.Rollback()
		return
	}
	var jsonData map[string]float32
	if err := ctx.ShouldBindJSON(&jsonData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}
	senderAccountId := float32(jsonData["senderAccountId"])
	recieverAccountId := float32(jsonData["recieverAccountId"])
	amount := float32(jsonData["amount"])

	_, err := queries.Debit(senderAccountId, amount, tx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}
	_, err2 := queries.Credit(recieverAccountId, amount, tx)
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		tx.Rollback()
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{"Amount transferred successfully": ""})
}
