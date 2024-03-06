package routes

import (
	"bankassignment/controllers"

	"github.com/gin-gonic/gin"
)

func BranchRoutes(router *gin.Engine) {
	branchroutes := router.Group("branch")
	branchroutes.POST("/", controllers.CreateBranch)
	branchroutes.DELETE("/:id", controllers.DeleteBranchById)
	branchroutes.GET("/:id", controllers.GetBranchDetailWithAllAccountAndCustomerDetails)
	branchroutes.PATCH("/", controllers.UpdateBranch)
}
