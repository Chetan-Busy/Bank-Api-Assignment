package routes

import (
	"bankassignment/controllers"

	"github.com/gin-gonic/gin"
)

func AccountRoutes(router *gin.Engine) {
	accountroutes := router.Group("account")
	accountroutes.POST("/", controllers.CreateAccount)
	accountroutes.DELETE("/:id", controllers.DeleteAccountById)
	accountroutes.GET("/:id", controllers.GetAccountDetailsById)
	accountroutes.PATCH("/", controllers.UpdateAccount)
	accountroutes.PATCH("/credit", controllers.Credit)
	accountroutes.PATCH("/debit", controllers.Debit)
	accountroutes.PATCH("/transfer", controllers.Transfer)
}
