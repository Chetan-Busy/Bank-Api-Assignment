package controllers

import (
	"bankassignment/models"
	"bankassignment/queries"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCustomer(ctx *gin.Context) {
	var customer *models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rows_inserted, err := queries.CreateCustomer(customer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Number of rows inserted": rows_inserted})
}

func GetCustomerDetailsById(ctx *gin.Context) {
	id := ctx.Param("id")
	customer, err := queries.GetCustomerDetailsById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Customer Details are as follows": customer})
}

func DeleteCustomerById(ctx *gin.Context) {
	id := ctx.Param("id")
	rows_deleted, err := queries.DeleteCustomerById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Number of rows deleted": rows_deleted})
}

func UpdateCustomer(ctx *gin.Context) {
	var customer *models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rows_updated, err := queries.UpdateCustomer(customer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if rows_updated == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No row found with that Id"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Customer updated successfully! No of row updated : ": rows_updated})
}
