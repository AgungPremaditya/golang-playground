package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "API_STATUS: TEST")
	})

	MovieRoutes(router)
	CategoryRoutes(router)

	return router
}
