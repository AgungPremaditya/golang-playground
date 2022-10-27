package controllers

import (
	"movies-golang-api/helpers"
	"movies-golang-api/models"
	"movies-golang-api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory(ctx *gin.Context) {
	var category models.Category

	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := repository.CreateCategoryQuery(category)

	ctx.JSON(http.StatusCreated, result)
}

func IndexCategory(ctx *gin.Context) {
	result := repository.IndexCategoryQuery()

	ctx.JSON(http.StatusOK, result)
}

func GetCategory(ctx *gin.Context) {
	// Convert params id from string to int
	id := helpers.StrToInt(ctx.Param("id"))

	result := repository.GetCategoryQuery(id)

	// If result not found return 404
	if result.CategoryID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

func UpdateCategory(ctx *gin.Context) {
	// Bind body params to var movie
	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Convert string id param to int
	id := helpers.StrToInt(ctx.Param("id"))

	// Checking is there any data with these id
	checkingResult := repository.GetCategoryQuery(id)

	// If there isn't any data return 404
	if checkingResult.CategoryID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})
	} else {
		repository.UpdateCategoryQuery(id, category)

		ctx.JSON(http.StatusOK, category)
	}
}

func DeleteCategory(ctx *gin.Context) {
	// Convert string id param to int
	id := helpers.StrToInt(ctx.Param("id"))

	// Checking is there any data with these id
	checkingResult := repository.GetCategoryQuery(id)

	// If there isn't any data return 404
	if checkingResult.CategoryID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})
	} else {
		// If there is match data then update these data
		repository.DeleteCategoryQuery(id)

		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}
