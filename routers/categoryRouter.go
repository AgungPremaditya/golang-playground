package routers

import (
	"movies-golang-api/controllers"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	router.POST("/categories", controllers.CreateCategory)
	router.GET("/categories", controllers.IndexCategory)
	router.GET("/categories/:id", controllers.GetCategory)
	router.PUT("/categories/:id", controllers.UpdateCategory)
	router.DELETE("/categories/:id", controllers.DeleteCategory)
}
