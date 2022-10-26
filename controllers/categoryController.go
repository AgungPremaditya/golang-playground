package controllers

import (
	"database/sql"
	"fmt"
	"movies-golang-api/models"
	"movies-golang-api/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCategory(ctx *gin.Context) {
	var category models.Category

	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := repository.CreateCategoryQuery(category)

	ctx.JSON(http.StatusCreated, result)
}

func IndexCategory(ctx *gin.Context) {
	result := repository.IndexCategoryQuery()

	ctx.JSON(http.StatusOK, result)
}

func GetCategory(ctx *gin.Context) {

	// Convert params id from string to int
	id := ctx.Param("id")
	categoryId, error := strconv.Atoi(id)
	checkError(error, "Convert Failed")

	result := repository.GetCategoryQuery(categoryId)

	// If result not found return 404
	if result.CategoryID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

func UpdateCategory(ctx *gin.Context) {
	// Bind body params to var movie
	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Convert string id param to int
	id := ctx.Param("id")
	categoryId, error := strconv.Atoi(id)
	checkError(error, "Convert Failed")

	// Checking is there any data with these id
	var resultId int
	script := fmt.Sprintf(`SELECT id FROM categories WHERE id = %d`, categoryId)
	rows := initDb().QueryRow(script)
	err := rows.Scan(&resultId)

	// If there isn't any data return 404
	if err != nil && err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})
	} else {
		// If there is match data then update these data
		script := `UPDATE categories SET categories.category_name = ?, categories.details = ? WHERE id = ?`
		rows, err := initDb().Query(script, category.CategoryName, category.Details, categoryId)
		checkError(err, "Query Failed")
		defer rows.Close()
		category.CategoryID = int64(categoryId)

		ctx.JSON(http.StatusOK, category)
	}
}

func DeleteCategory(ctx *gin.Context) {
	// Convert string id param to int
	id := ctx.Param("id")
	categoryId, error := strconv.Atoi(id)
	checkError(error, "Convert Failed")

	// Checking is there any data with these id
	var resultId int
	script := fmt.Sprintf(`SELECT id FROM categories WHERE id = %d`, categoryId)
	rows := initDb().QueryRow(script)
	err := rows.Scan(&resultId)

	// If there isn't any data return 404
	if err != nil && err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})
	} else {
		// If there is match data then update these data
		script := `DELETE FROM categories WHERE id = ?`
		rows, err := initDb().Exec(script, categoryId)
		checkError(err, "Query Failed")
		defer rows.RowsAffected()

		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}
