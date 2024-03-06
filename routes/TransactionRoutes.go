package routes

import (
	"bankassignment/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(router *gin.Engine) {
	transactionRoutes := router.Group("transaction")
	transactionRoutes.POST("/", controllers.CreateTransaction)
	transactionRoutes.GET("/account/:id", controllers.GetTransactionByAccount)
	transactionRoutes.GET("/:id", controllers.GetTransactionDetailsById)
}
