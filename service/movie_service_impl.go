package service

import (
	"context"
	"database/sql"
	"movies-golang-api/helpers"
	"movies-golang-api/models/domain"
	"movies-golang-api/models/web"
	"movies-golang-api/repository"

	"github.com/go-playground/validator/v10"
)

type MovieServiceImpl struct {
	MovieRepository repository.MovieRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewMovieService(movieRepository repository.MovieRepository, DB *sql.DB, validate *validator.Validate) MovieService {
	return &MovieServiceImpl{
		MovieRepository: movieRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *MovieServiceImpl) Create(ctx context.Context, request web.MovieCreateRequest) web.MovieResponse {
	err := service.Validate.Struct(request)
	helpers.CheckError(err)

	tx, err := service.DB.Begin()
	helpers.CheckError(err)
	defer helpers.CommitOrRollback(tx)

	movie := domain.Movie{
		Title:      request.Title,
		Rating:     request.Rating,
		Details:    request.Details,
		CategoryId: request.CategoryId,
	}

	movie = service.MovieRepository.Save(ctx, tx, movie)

	return helpers.ToMovieResponse(movie)
}

func (service *MovieServiceImpl) Update(ctx context.Context, request web.MovieUpdateRequest) web.MovieResponse {
	err := service.Validate.Struct(request)
	helpers.CheckError(err)

	tx, err := service.DB.Begin()
	helpers.CheckError(err)
	defer helpers.CommitOrRollback(tx)

	movie, err := service.MovieRepository.FindById(ctx, tx, request.Id)
	helpers.CheckError(err)

	movie.Title = request.Title
	movie.Rating = request.Rating
	movie.Details = request.Details
	movie.CategoryId = request.CategoryId

	result := service.MovieRepository.Update(ctx, tx, helpers.ToMovieDomain(movie))

	return helpers.ToMovieResponse(result)
}

func (service *MovieServiceImpl) Delete(ctx context.Context, movieId int) {
	tx, err := service.DB.Begin()
	helpers.CheckError(err)
	defer helpers.CommitOrRollback(tx)

	movie, err := service.MovieRepository.FindById(ctx, tx, movieId)
	helpers.CheckError(err)

	service.MovieRepository.Delete(ctx, tx, helpers.ToMovieDomain(movie))
}

func (service *MovieServiceImpl) FindById(ctx context.Context, movieId int) web.MovieWithCategoryResponse {
	tx, err := service.DB.Begin()
	helpers.CheckError(err)
	defer helpers.CommitOrRollback(tx)

	movie, err := service.MovieRepository.FindById(ctx, tx, movieId)
	helpers.CheckError(err)

	return web.MovieWithCategoryResponse(movie)
}

func (service *MovieServiceImpl) FindAll(ctx context.Context) []web.MovieWithCategoryResponse {
	tx, err := service.DB.Begin()
	helpers.CheckError(err)
	defer helpers.CommitOrRollback(tx)

	movies := service.MovieRepository.FindAll(ctx, tx)

	var movieResponses []web.MovieWithCategoryResponse
	for _, movie := range movies {
		movieResponses = append(movieResponses, helpers.ToMovieWithCategoryResponse(movie))
	}

	return movieResponses
}
