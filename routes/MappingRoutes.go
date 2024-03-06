package routes

import (
	"bankassignment/controllers"

	"github.com/gin-gonic/gin"
)

func MappingRoutes(router *gin.Engine) {
	mappingroutes := router.Group("mapping")
	mappingroutes.POST("/", controllers.CreateMapping)
	mappingroutes.DELETE("/:id", controllers.DeleteMapping)
	mappingroutes.PATCH("/", controllers.UpdateMapping)
}
