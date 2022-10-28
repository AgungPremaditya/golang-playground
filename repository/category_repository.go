package repository

import (
	"context"
	"database/sql"
	"movies-golang-api/models"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category models.Category) models.Category
	Update(ctx context.Context, tx *sql.Tx, category models.Category) models.Category
	Delete(ctx context.Context, tx *sql.Tx, category models.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) models.Category
	FindAll(ctx context.Context, tx *sql.Tx) []models.Category
}
