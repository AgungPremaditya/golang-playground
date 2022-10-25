package controllers

import (
	"log"
	"movies-golang-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func CreateMovie(ctx *gin.Context) {
	var movie models.Movie

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	script := `INSERT INTO movie (title, rating, details) VALUES (?, ?, ?)`

	rows, err := initDb().Exec(script, movie.Title, movie.Rating, movie.Description)
	checkError(err, "Insert Failed")
	lastId, _ := rows.LastInsertId()
	movie.MovieID = lastId

	ctx.JSON(http.StatusCreated, movie)
}

func IndexMovie(ctx *gin.Context) {
	var movies models.Movie
	var result []models.Movie

	script := `SELECT * FROM movie`
	rows, err := initDb().Query(script)
	checkError(err, "Get Index Failed")
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&movies.MovieID, &movies.Title, &movies.Rating, &movies.Description)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, movies)
	}

	ctx.JSON(http.StatusOK, result)
}

func GetMovie(ctx *gin.Context) {
	var movie models.Movie
	id := ctx.Param("id")

	script := `SELECT * FROM movie WHERE id = ?`
	rows, err := initDb().Query(script, id)
	checkError(err, "Get Index Failed")
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&movie.MovieID, &movie.Title, &movie.Rating, &movie.Description)
		if err != nil {
			log.Fatal(err)
		}
	}

	ctx.JSON(http.StatusOK, movie)
}

func UpdateMovie(ctx *gin.Context) {
	var movie models.Movie
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	script := `UPDATE movie SET movie.title = ?, movie.rating = ?, movie.details = ? WHERE id = ?`
	rows, err := initDb().Query(script, movie.Title, movie.Rating, movie.Description, id)
	checkError(err, "Get Index Failed")
	defer rows.Close()

	movieId, err := strconv.Atoi(id)
	checkError(err, "Convert Failed")
	movie.MovieID = int64(movieId)

	ctx.JSON(http.StatusOK, movie)
}

// func DeleteMovie(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	condition := false
// 	var index int

// 	for i, movie := range movies {
// 		if id == movie.MovieID {
// 			condition = true
// 			index = i
// 			break
// 		}
// 	}

// 	if !condition {
// 		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 			"error": "Data not found",
// 		})
// 		return
// 	}

// 	copy(movies[index:], movies[index+1:])
// 	movies[len(movies)-1] = Movie{}
// 	movies = movies[:len(movies)-1]

// 	ctx.JSON(http.StatusNoContent, gin.H{})
// }
