package routers

import (
	"movies-golang-api/controllers"
	"movies-golang-api/database"
	"movies-golang-api/repository"
	"movies-golang-api/service"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {

	db := database.InitDb()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db)
	categoryController := controllers.NewCategoryController(categoryService)

	categories := router.Group("/categories")
	{
		categories.POST("/", categoryController.Create)
		categories.GET("/", categoryController.FindAll)
		categories.GET("/:id", categoryController.FindById)
		categories.PUT("/:id", categoryController.Update)
		categories.DELETE("/:id", categoryController.Delete)
	}
}
