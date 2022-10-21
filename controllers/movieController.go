package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	MovieID     string `json:"movie_id"`
	Title       string `json:"title"`
	Rating      int    `json:"rating"`
	Description string `json:"desc"`
}

var movies = []Movie{}

func CreateMovie(ctx *gin.Context) {
	var newMovie Movie

	if err := ctx.ShouldBindJSON(&newMovie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newMovie.MovieID = fmt.Sprintf("c%d", len(movies)+1)
	movies = append(movies, newMovie)

	ctx.JSON(http.StatusCreated, newMovie)
}

func IndexMovie(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, movies)
}

func GetMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	condition := false
	var movieData Movie

	for i, movie := range movies {
		if movie.MovieID == id {
			condition = true
			movieData = movies[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, movieData)
}

func UpdateMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	condition := false
	var movieData Movie

	if err := ctx.ShouldBindJSON(&movieData); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, movie := range movies {
		if id == movie.MovieID {
			condition = true
			movies[i] = movieData
			movies[i].MovieID = id
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, movieData)
}

func DeleteMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	condition := false
	var index int

	for i, movie := range movies {
		if id == movie.MovieID {
			condition = true
			index = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Data not found",
		})
		return
	}

	copy(movies[index:], movies[index+1:])
	movies[len(movies)-1] = Movie{}
	movies = movies[:len(movies)-1]

	ctx.JSON(http.StatusNoContent, gin.H{})
}
