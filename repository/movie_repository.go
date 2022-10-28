package repository

import (
	"context"
	"database/sql"
	"movies-golang-api/models/domain"
)

type MovieRepository interface {
	Save(ctx context.Context, tx *sql.Tx, movie domain.Movie) domain.Movie
	Update(ctx context.Context, tx *sql.Tx, movie domain.Movie) domain.Movie
	Delete(ctx context.Context, tx *sql.Tx, movie domain.Movie)
	FindById(ctx context.Context, tx *sql.Tx, movieId int) (domain.MovieWithCategory, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.MovieWithCategory
}
