package repository

import (
	"context"
	"database/sql"
	"errors"
	"movies-golang-api/helpers"
	"movies-golang-api/models/domain"
)

type MovieRepositoryImpl struct {
}

func (repository *MovieRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, movie domain.Movie) domain.Movie {
	query := "INSERT INTO movies (title, rating, details, category_id) VALUES (?, ?, ?, ?)"

	result, err := tx.ExecContext(ctx, query, movie.Title, movie.Rating, movie.Details, movie.CategoryId)
	helpers.CheckError(err)

	id, err := result.LastInsertId()
	helpers.CheckError(err)

	movie.Id = int(id)

	return movie
}

func (repository *MovieRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, movie domain.Movie) domain.Movie {
	query := "UPDATE movies SET movies.title = ?, movies.rating = ?, movies.details = ?, movies.category_id = ? WHERE movies.id = ?"
	_, err := tx.ExecContext(ctx, query, movie.Title, movie.Rating, movie.Details, movie.CategoryId)
	helpers.CheckError(err)

	return movie
}

func (repository *MovieRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, movie domain.Movie) {
	query := "DELETE FROM movies WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, movie.Id)
	helpers.CheckError(err)
}

func (repository *MovieRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, movieId int) (domain.MovieWithCategory, error) {
	query := "SELECT movies.id, movies.title, movies.rating, movies.details, categories.id as category_id, categories.category_name, categories.details as category_details FROM movies INNER JOIN categories ON movies.category_id = categories.id WHERE movies.id= ?"
	rows, err := tx.QueryContext(ctx, query, movieId)
	helpers.CheckError(err)
	defer rows.Close()

	movie := domain.MovieWithCategory{}
	if rows.Next() {
		rows.Scan(&movie.Id, &movie.Title, &movie.Rating, &movie.Details, &movie.CategoryId, &movie.CategoryName, &movie.CategoryDetails)
		return movie, nil
	} else {
		return movie, errors.New("movie not found")
	}
}

func (repository *MovieRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.MovieWithCategory {
	query := "SELECT movies.id, movies.title, movies.rating, movies.details, categories.id as category_id, categories.category_name, categories.details as category_details FROM movies INNER JOIN categories ON movies.category_id = categories.id"
	rows, err := tx.Query(query)
	helpers.CheckError(err)
	defer rows.Close()

	var movies []domain.MovieWithCategory
	for rows.Next() {
		movie := domain.MovieWithCategory{}
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Rating, &movie.Details, &movie.CategoryId, &movie.CategoryName, &movie.CategoryDetails)
		helpers.CheckError(err)
		movies = append(movies, movie)
	}

	return movies
}
