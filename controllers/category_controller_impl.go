package controllers

import (
	"movies-golang-api/helpers"
	"movies-golang-api/models/web"
	"movies-golang-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(ctx *gin.Context) {
	var body web.CategoryCreateRequest

	err := ctx.ShouldBind(&body)
	helpers.CheckError(err)

	result := controller.CategoryService.Create(ctx, body)

	ctx.JSON(http.StatusOK, result)
}

func (controller *CategoryControllerImpl) Update(ctx *gin.Context) {
	var body web.CategoryUpdateRequest

	err := ctx.ShouldBind(&body)
	helpers.CheckError(err)

	categoryId, err := strconv.Atoi(ctx.Param("id"))
	helpers.CheckError(err)

	body.Id = categoryId

	result := controller.CategoryService.Update(ctx, body)

	ctx.JSON(http.StatusOK, result)
}

func (controller *CategoryControllerImpl) Delete(ctx *gin.Context) {
	categoryId, err := strconv.Atoi(ctx.Param("id"))
	helpers.CheckError(err)

	controller.CategoryService.Delete(ctx, categoryId)

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (controller *CategoryControllerImpl) FindById(ctx *gin.Context) {
	categoryId, err := strconv.Atoi(ctx.Param("id"))
	helpers.CheckError(err)

	result := controller.CategoryService.FindById(ctx, categoryId)

	ctx.JSON(http.StatusOK, result)
}

func (controller *CategoryControllerImpl) FindAll(ctx *gin.Context) {
	result := controller.CategoryService.FindAll(ctx)

	ctx.JSON(http.StatusOK, result)
}
