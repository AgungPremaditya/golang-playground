package helpers

import (
	"movies-golang-api/models/domain"
	"movies-golang-api/models/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:      category.Id,
		Name:    category.Name,
		Details: category.Details,
	}
}

func ToMovieResponse(movie domain.Movie) web.MovieResponse {
	return web.MovieResponse{
		Id:         movie.Id,
		Title:      movie.Title,
		Rating:     movie.Rating,
		Details:    movie.Details,
		CategoryId: movie.CategoryId,
	}
}

func ToMovieWithCategoryResponse(movie domain.MovieWithCategory) web.MovieWithCategoryResponse {
	return web.MovieWithCategoryResponse{
		Id:              movie.Id,
		Title:           movie.Title,
		Rating:          movie.Rating,
		Details:         movie.Details,
		CategoryId:      movie.CategoryId,
		CategoryName:    movie.CategoryName,
		CategoryDetails: movie.CategoryDetails,
	}
}

func ToMovieDomain(movie domain.MovieWithCategory) domain.Movie {
	return domain.Movie{
		Id:         movie.Id,
		Title:      movie.Title,
		Rating:     movie.Rating,
		Details:    movie.Details,
		CategoryId: movie.CategoryId,
	}
}
