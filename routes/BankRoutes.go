package routes

import (
	"bankassignment/controllers"

	"github.com/gin-gonic/gin"
)

func BankRoutes(router *gin.Engine) {
	bankRoutes := router.Group("bank")
	bankRoutes.POST("/", controllers.CreateBank)
	bankRoutes.GET("/", controllers.GetAllBanks)
	bankRoutes.PATCH("/", controllers.UpdateBank)
	bankRoutes.DELETE("/", controllers.DeleteAllBanks)
	bankRoutes.DELETE("/:id", controllers.DeleteBankById)
	bankRoutes.GET("/branch", controllers.GetAllBanksWithBranches)
	bankRoutes.GET("/:id", controllers.GetBankDetailById)
}
