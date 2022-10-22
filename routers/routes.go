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

	return router
}

// migrate -database "postgres://postgres:241403@localhost:5432/golang_movie_api" -path db/migrations up
