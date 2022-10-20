package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	MovieID     string `json:"car_id"`
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
