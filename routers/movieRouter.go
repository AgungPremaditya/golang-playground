package routers

import (
	"movies-golang-api/controllers"

	"github.com/gin-gonic/gin"
)

// func StartServer() *gin.Engine {
// 	router := gin.Default()

// 	router.POST("/movies", controllers.CreateMovie)
// 	router.GET("/movies", controllers.IndexMovie)
// 	return router
// }

func MovieRoutes(router *gin.Engine) {
	router.POST("/movies", controllers.CreateMovie)
	router.GET("/movies", controllers.IndexMovie)
}
