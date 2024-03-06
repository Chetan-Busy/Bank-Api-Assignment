package controllers

import (
	"bankassignment/models"
	"bankassignment/queries"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBank(ctx *gin.Context) {
	var bank models.Bank
	if err := ctx.ShouldBindJSON(&bank); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := queries.CreateBank(&bank)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"Bank created with id: ": id})
}

func GetAllBanks(ctx *gin.Context) {
	banks, err := queries.GetAllBanks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"Banks: ": banks})
}

func UpdateBank(ctx *gin.Context) {
	var bank *models.Bank
	if err := ctx.ShouldBindJSON(&bank); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	res, err := queries.UpdateBank(bank)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if res == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Update successfull. No of rows updated": res})

}

func DeleteAllBanks(ctx *gin.Context) {
	rows_affected, err := queries.DeleteAllBanks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"Total number of rows deleted": rows_affected})
}

func DeleteBankById(ctx *gin.Context) {
	id := ctx.Param("id")
	rows_affected, err := queries.DeleteBankById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if rows_affected == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No row with this id found!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Deleted Successfully, No of rows deleted : ": rows_affected})
}

func GetAllBanksWithBranches(ctx *gin.Context) {
	banks, err := queries.GetAllBanksWithBranches()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"All banks with branches are : ": banks})
}

func GetBankDetailById(ctx *gin.Context) {
	id := ctx.Param("id")
	bank, err := queries.GetBankDetailsById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Bank details": bank})

}
