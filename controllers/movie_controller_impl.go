package controllers

import (
	"movies-golang-api/helpers"
	"movies-golang-api/models/web"
	"movies-golang-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieControllerImpl struct {
	MovieService service.MovieService
}

func NewMovieController(movieService service.MovieService) MovieController {
	return &MovieControllerImpl{
		MovieService: movieService,
	}
}

func (controller *MovieControllerImpl) Create(ctx *gin.Context) {
	var body web.MovieCreateRequest

	err := ctx.ShouldBind(&body)
	helpers.CheckError(err)

	result := controller.MovieService.Create(ctx, body)

	ctx.JSON(http.StatusOK, result)
}

func (controller *MovieControllerImpl) Update(ctx *gin.Context) {
	var body web.MovieUpdateRequest

	err := ctx.ShouldBind(&body)
	helpers.CheckError(err)

	movieId, err := strconv.Atoi(ctx.Param("id"))
	helpers.CheckError(err)

	body.Id = movieId

	result := controller.MovieService.Update(ctx, body)

	ctx.JSON(http.StatusOK, result)
}

func (controller *MovieControllerImpl) Delete(ctx *gin.Context) {
	movieId, err := strconv.Atoi(ctx.Param("id"))
	helpers.CheckError(err)

	controller.MovieService.Delete(ctx, movieId)

	ctx.JSON(http.StatusNoContent, gin.H{})

}

func (controller *MovieControllerImpl) FindById(ctx *gin.Context) {
	movieId, err := strconv.Atoi(ctx.Param("id"))
	helpers.CheckError(err)

	result := controller.MovieService.FindById(ctx, movieId)

	ctx.JSON(http.StatusOK, result)
}

func (controller *MovieControllerImpl) FindAll(ctx *gin.Context) {
	result := controller.MovieService.FindAll(ctx)

	ctx.JSON(http.StatusOK, result)
}
