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
