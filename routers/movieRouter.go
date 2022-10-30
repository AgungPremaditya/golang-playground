package routers

import (
	"movies-golang-api/controllers"
	"movies-golang-api/database"
	"movies-golang-api/repository"
	"movies-golang-api/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func MovieRoutes(router *gin.Engine) {
	validate := validator.New()
	db := database.InitDb()

	movieRepository := repository.NewMovieRepository()
	movieService := service.NewMovieService(movieRepository, db, validate)
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
