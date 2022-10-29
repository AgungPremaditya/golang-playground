package routers

import (
	"movies-golang-api/controllers"
	"movies-golang-api/database"
	"movies-golang-api/repository"
	"movies-golang-api/service"

	"github.com/gin-gonic/gin"
)

func MovieRoutes(router *gin.Engine) {
	db := database.InitDb()
	movieRepository := repository.NewMovieRepository()
	movieService := service.NewMovieService(movieRepository, db)
	movieController := controllers.NewMovieController(movieService)

	movies := router.Group("/movies")
	{
		movies.POST("/", movieController.Create)
		movies.GET("/", movieController.FindAll)
		movies.GET("/:id", movieController.FindById)
		movies.PUT("/:id", movieController.Update)
		movies.DELETE("/:id", movieController.Delete)
	}
}
