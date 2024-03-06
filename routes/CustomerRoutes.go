package routes

import (
	"bankassignment/controllers"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(router *gin.Engine) {
	customerRoutes := router.Group("customer")
	customerRoutes.POST("/", controllers.CreateCustomer)
	customerRoutes.GET("/:id", controllers.GetCustomerDetailsById)
	customerRoutes.DELETE("/:id", controllers.DeleteCustomerById)
	customerRoutes.PATCH("/", controllers.UpdateCustomer)
}
