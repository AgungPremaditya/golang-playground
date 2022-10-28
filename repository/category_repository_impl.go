package repository

import (
	"context"
	"database/sql"
	"errors"
	"movies-golang-api/helpers"
	"movies-golang-api/models/domain"
)

type CategoryRepositoryImpl struct {
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "INSERT INTO categories (category_name, details) VALUES (?, ?)"

	result, err := tx.ExecContext(ctx, query, category.Name, category.Details)
	helpers.CheckError(err)

	id, err := result.LastInsertId()
	helpers.CheckError(err)

	category.Id = int(id)

	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "UPDATE categories SET categories.category_name = ?, categories.details = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, category.Name, category.Details, category.Id)
	helpers.CheckError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	query := "DELETE FROM categories WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, category.Id)
	helpers.CheckError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	query := "SELECT * FROM categories WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, categoryId)
	helpers.CheckError(err)

	category := domain.Category{}
	if rows.Next() {
		rows.Scan(&category.Id, &category.Name, &category.Details)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	query := `SELECT * FROM categories`
	rows, err := tx.QueryContext(ctx, query)
	helpers.CheckError(err)

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name, &category.Details)
		helpers.CheckError(err)
		categories = append(categories, category)
	}

	return categories
}
