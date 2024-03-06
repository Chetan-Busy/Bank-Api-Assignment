package controllers

import (
	"bankassignment/models"
	"bankassignment/queries"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBranch(ctx *gin.Context) {
	var branch *models.Branch
	if err := ctx.ShouldBindJSON(&branch); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rows_inserted, err := queries.CreateBranch(branch)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"No. of rows inserted: ": rows_inserted})
}

func DeleteBranchById(ctx *gin.Context) {
	id := ctx.Param("id")
	rows_affected, err := queries.DeleteBranchById(id)
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

func GetBranchDetailWithAllAccountAndCustomerDetails(ctx *gin.Context) {
	id := ctx.Param("id")
	branch, err := queries.GetBranchDetailWithAllAccountAndCustomerDetails(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Branch Details:": branch})

}

func UpdateBranch(ctx *gin.Context) {
	var branch *models.Branch
	if err := ctx.ShouldBindJSON(&branch); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rows_updated, err := queries.UpdateBranch(branch)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if rows_updated == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No branch found with that ID"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Branch updated successfully. No. of rows updated ": rows_updated})
}
