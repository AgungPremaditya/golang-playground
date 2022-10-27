package repository

import (
	"database/sql"
	"fmt"
	"log"
	"movies-golang-api/database"
	"movies-golang-api/models"
)

func CreateMovieQuery(movie models.Movie) models.Movie {
	var result models.Movie

	script := `INSERT INTO movies (title, rating, details, category_id) VALUES (?, ?, ?, ?)`

	rows, err := database.InitDb().Exec(script, movie.Title, movie.Rating, movie.Description, movie.CategoryID)
	database.CheckError(err, "Insert Failed")
	lastId, _ := rows.LastInsertId()

	result.MovieID = lastId
	result.Title = movie.Title
	result.Rating = movie.Rating
	result.Description = movie.Description
	result.CategoryID = movie.CategoryID

	return result
}

func IndexMovieQuery() []models.MovieWithCategory {
	var movies models.MovieWithCategory
	var result []models.MovieWithCategory

	script := `SELECT movies.id, movies.title, movies.rating, movies.details, categories.id as category_id, categories.category_name, categories.details as category_details FROM movies INNER JOIN categories ON movies.category_id = categories.id`
	rows, err := database.InitDb().Query(script)
	database.CheckError(err, "Get Index Failed")
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&movies.MovieID, &movies.Title, &movies.Rating, &movies.Description, &movies.CategoryID, &movies.CategoryName, &movies.CategoryDetails)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, movies)
	}

	return result
}

func GetMovieQuery(id int) models.MovieWithCategory {
	var result models.MovieWithCategory

	// Find data with id
	script := fmt.Sprintf(`SELECT movies.id, movies.title, movies.rating, movies.details, categories.id as category_id, categories.category_name, categories.details as category_details FROM movies INNER JOIN categories ON movies.category_id = categories.id WHERE movies.id=%d`, id)
	rows := database.InitDb().QueryRow(script)
	err := rows.Scan(&result.MovieID, &result.Title, &result.Rating, &result.Description, &result.CategoryID, &result.CategoryName, &result.CategoryDetails)

	if err != nil && err == sql.ErrNoRows {
		return result
	} else {
		database.CheckError(err, "Get Data Failed")
	}

	return result
}

func UpdateMovieQuery(id int, movie models.Movie) models.Movie {
	var result models.Movie

	script := `UPDATE movies SET movies.title = ?, movies.rating = ?, movies.details = ?, movies.category_id = ? WHERE movies.id = ?`
	rows, err := database.InitDb().Exec(script, movie.Title, movie.Rating, movie.Description, movie.CategoryID, id)
	database.CheckError(err, "Query Failed")
	affectedRow, _ := rows.RowsAffected()

	result.MovieID = int64(affectedRow)
	result.Title = movie.Title
	result.Rating = movie.Rating
	result.Description = movie.Description
	result.CategoryID = movie.CategoryID

	return result
}

func DeleteMovieQuery(id int) int64 {
	script := `DELETE FROM movies WHERE id = ?`
	rows, err := database.InitDb().Exec(script, id)
	database.CheckError(err, "Query Failed")
	AffectedRowId, _ := rows.RowsAffected()

	return AffectedRowId
}
