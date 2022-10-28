package service

import (
	"context"
	"database/sql"
	"movies-golang-api/helpers"
	"movies-golang-api/models/domain"
	"movies-golang-api/models/web"
	"movies-golang-api/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helpers.CheckError(err)
	defer helpers.CommitOrRollback(tx)

	category := domain.Category{
		Name:    request.Name,
		Details: request.Details,
	}

	service.CategoryRepository.Save(ctx, tx, category)

	return helpers.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helpers.CheckError(err)
	defer helpers.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helpers.CheckError(err)

	category.Name = request.Name
	category.Details = request.Details

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helpers.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helpers.CheckError(err)
	defer helpers.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helpers.CheckError(err)

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helpers.CheckError(err)
	defer helpers.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helpers.CheckError(err)

	return web.CategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helpers.CheckError(err)
	defer helpers.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, helpers.ToCategoryResponse(category))
	}

	return categoryResponses
}
