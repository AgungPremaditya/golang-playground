package routers

import (
	"movies-golang-api/controllers"

	"github.com/gin-gonic/gin"
)

func MovieRoutes(router *gin.Engine) {
	router.POST("/movies", controllers.CreateMovie)
	router.GET("/movies", controllers.IndexMovie)
	router.GET("/movies/:id", controllers.GetMovie)
}
