package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"movies-golang-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type MovieWithCategory struct {
	models.Movie
	CategoryName    string `json:"category_name"`
	CategoryDetails string `json:"category_details"`
}

func CreateMovie(ctx *gin.Context) {
	var movie models.Movie

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	script := `INSERT INTO movies (title, rating, details, category_id) VALUES (?, ?, ?, ?)`

	rows, err := initDb().Exec(script, movie.Title, movie.Rating, movie.Description, movie.CategoryID)
	checkError(err, "Insert Failed")
	lastId, _ := rows.LastInsertId()
	movie.MovieID = lastId

	ctx.JSON(http.StatusCreated, movie)
}

func IndexMovie(ctx *gin.Context) {
	var movies MovieWithCategory
	var result []MovieWithCategory

	script := `SELECT movies.id, movies.title, movies.rating, movies.details, categories.id as category_id, categories.category_name, categories.details as category_details FROM movies INNER JOIN categories ON movies.category_id = categories.id`
	rows, err := initDb().Query(script)
	checkError(err, "Get Index Failed")
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&movies.MovieID, &movies.Title, &movies.Rating, &movies.Description, &movies.CategoryID, &movies.CategoryName, &movies.CategoryDetails)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, movies)
	}

	ctx.JSON(http.StatusOK, result)
}

func GetMovie(ctx *gin.Context) {
	var movie MovieWithCategory

	// Convert params id from string to int
	id := ctx.Param("id")
	movieId, error := strconv.Atoi(id)
	checkError(error, "Convert Failed")

	// Find data with id
	script := fmt.Sprintf(`SELECT movies.id, movies.title, movies.rating, movies.details, categories.id as category_id, categories.category_name, categories.details as category_details FROM movies INNER JOIN categories ON movies.category_id = categories.id WHERE movies.id=%d`, movieId)
	rows := initDb().QueryRow(script)
	err := rows.Scan(&movie.MovieID, &movie.Title, &movie.Rating, &movie.Description, &movie.CategoryID, &movie.CategoryName, &movie.CategoryDetails)

	// If there isn't match id return 404
	if err != nil && err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})
	} else {
		// If there is match id return data
		ctx.JSON(http.StatusOK, movie)
	}
}

func UpdateMovie(ctx *gin.Context) {
	// Bind body params to var movie
	var movie models.Movie
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Convert string id param to int
	id := ctx.Param("id")
	movieId, error := strconv.Atoi(id)
	checkError(error, "Convert Failed")

	// Checking is there any data with these id
	var resultId int
	script := fmt.Sprintf(`SELECT id FROM movies WHERE id = %d`, movieId)
	rows := initDb().QueryRow(script)
	err := rows.Scan(&resultId)

	// If there isn't any data return 404
	if err != nil && err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})
	} else {
		// If there is match data then update these data
		script := `UPDATE movies SET movies.title = ?, movies.rating = ?, movies.details = ?, movies.category_id = ? WHERE movies.id = ?`
		rows, err := initDb().Query(script, movie.Title, movie.Rating, movie.Description, movie.CategoryID, movieId)
		checkError(err, "Query Failed")
		defer rows.Close()
		movie.MovieID = int64(movieId)

		ctx.JSON(http.StatusOK, movie)
	}
}

func DeleteMovie(ctx *gin.Context) {
	// Convert string id param to int
	id := ctx.Param("id")
	movieId, error := strconv.Atoi(id)
	checkError(error, "Convert Failed")

	// Checking is there any data with these id
	var resultId int
	script := fmt.Sprintf(`SELECT id FROM movies WHERE id = %d`, movieId)
	rows := initDb().QueryRow(script)
	err := rows.Scan(&resultId)

	// If there isn't any data return 404
	if err != nil && err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})
	} else {
		// If there is match data then update these data
		script := `DELETE FROM movies WHERE id = ?`
		rows, err := initDb().Exec(script, movieId)
		checkError(err, "Query Failed")
		defer rows.RowsAffected()

		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}
