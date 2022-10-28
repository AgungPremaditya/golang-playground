package service

import (
	"context"
	"movies-golang-api/models/web"
)

type MovieService interface {
	Create(ctx context.Context, request web.MovieCreateRequest) web.MovieResponse
	Update(ctx context.Context, request web.MovieUpdateRequest) web.MovieResponse
	Delete(ctx context.Context, movieId int)
	FindById(ctx context.Context, movieId int) web.MovieWithCategoryResponse
	FindAll(ctx context.Context) []web.MovieWithCategoryResponse
}
