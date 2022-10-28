package repository

import (
	"context"
	"database/sql"
	"errors"
	"movies-golang-api/helpers"
	"movies-golang-api/models"
)

type CategoryRepositoryImpl struct {
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category models.Category) models.Category {
	query := "INSERT INTO categories (category_name, details) VALUES (?, ?)"

	result, err := tx.ExecContext(ctx, query, category.CategoryName, category.Details)
	helpers.CheckError(err)

	id, err := result.LastInsertId()
	helpers.CheckError(err)

	category.CategoryID = int(id)

	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category models.Category) models.Category {
	query := "UPDATE categories SET categories.category_name = ?, categories.details = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, category.CategoryName, category.Details, category.CategoryID)
	helpers.CheckError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category models.Category) {
	query := "DELETE FROM categories WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, category.CategoryID)
	helpers.CheckError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (models.Category, error) {
	query := "SELECT * FROM categories WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, categoryId)
	helpers.CheckError(err)

	category := models.Category{}
	if rows.Next() {
		rows.Scan(&category.CategoryID, &category.CategoryName, &category.Details)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []models.Category {
	query := `SELECT * FROM categories`
	rows, err := tx.QueryContext(ctx, query)
	helpers.CheckError(err)

	var categories []models.Category
	for rows.Next() {
		category := models.Category{}
		err := rows.Scan(&category.CategoryID, &category.CategoryName, &category.Details)
		helpers.CheckError(err)
		categories = append(categories, category)
	}

	return categories
}
