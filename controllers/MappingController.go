package controllers

import (
	db "bankassignment/database"
	"bankassignment/models"
	"bankassignment/queries"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMapping(ctx *gin.Context) {
	tx, txErr := db.DB.Begin()
	if txErr != nil {
		log.Println("Error starting a transaction")
		return
	}

	var mapping *models.Mapping
	if err := ctx.ShouldBindJSON(&mapping); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	rows_inserted, err := queries.CreateMapping(mapping, mapping.CustomerId, mapping.AccountId, tx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{"Mapping inserted successfully. No. of rows inserted": rows_inserted})
}

func DeleteMapping(ctx *gin.Context) {
	id := ctx.Param("id")
	rows_deleted, err := queries.DeleteMapping(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if rows_deleted == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No. mapping with this id exists"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Mapping Deleted Successfully! Number of rows deleted": rows_deleted})
}

func UpdateMapping(ctx *gin.Context) {
	var mapping *models.Mapping
	if err := ctx.ShouldBindJSON(&mapping); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rows_updated, err := queries.UpdateMapping(mapping)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if rows_updated == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No mapping exits with given id"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Mapping updated successfully.Number of rows updated": rows_updated})
}
