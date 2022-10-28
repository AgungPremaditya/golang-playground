package controllers

// import (
// 	"movies-golang-api/helpers"
// 	"movies-golang-api/models"
// 	"movies-golang-api/repository"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/go-sql-driver/mysql"
// )

// func CreateMovie(ctx *gin.Context) {
// 	var movie models.Movie

// 	if err := ctx.ShouldBindJSON(&movie); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	result := repository.CreateMovieQuery(movie)

// 	ctx.JSON(http.StatusCreated, result)
// }

// func IndexMovie(ctx *gin.Context) {
// 	result := repository.IndexMovieQuery()

// 	ctx.JSON(http.StatusOK, result)
// }

// func GetMovie(ctx *gin.Context) {
// 	// Convert params id from string to int
// 	id := helpers.StrToInt(ctx.Param("id"))

// 	result := repository.GetMovieQuery(id)

// 	// If there isn't match id return 404
// 	if result.MovieID == 0 {
// 		ctx.JSON(http.StatusNotFound, gin.H{
// 			"error": "data not found",
// 		})
// 	} else {
// 		// If there is match id return data
// 		ctx.JSON(http.StatusOK, result)
// 	}
// }

// func UpdateMovie(ctx *gin.Context) {
// 	// Bind body params to var movie
// 	var movie models.Movie
// 	if err := ctx.ShouldBindJSON(&movie); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	// Convert string id param to int
// 	id := helpers.StrToInt(ctx.Param("id"))

// 	// Checking is there any data with these id
// 	checkingResult := repository.GetMovieQuery(id)

// 	// If there isn't any data return 404
// 	if checkingResult.MovieID == 0 {
// 		ctx.JSON(http.StatusNotFound, gin.H{
// 			"error": "data not found",
// 		})
// 	} else {
// 		result := repository.UpdateMovieQuery(id, movie)

// 		ctx.JSON(http.StatusOK, result)
// 	}
// }

// func DeleteMovie(ctx *gin.Context) {
// 	// Convert string id param to int
// 	id := helpers.StrToInt(ctx.Param("id"))

// 	// Checking is there any data with these id
// 	checkingResult := repository.GetMovieQuery(id)

// 	// If there isn't any data return 404
// 	if checkingResult.MovieID == 0 {
// 		ctx.JSON(http.StatusNotFound, gin.H{
// 			"error": "data not found",
// 		})
// 	} else {
// 		repository.DeleteMovieQuery(id)
// 		ctx.JSON(http.StatusNoContent, gin.H{})
// 	}
// }
