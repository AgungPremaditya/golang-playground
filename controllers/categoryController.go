package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"movies-golang-api/models"
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

	script := `INSERT INTO categories (category_name, details) VALUES (?, ?)`

	rows, err := initDb().Exec(script, category.CategoryName, category.Details)
	checkError(err, "Insert Failed")
	lastId, _ := rows.LastInsertId()
	category.CategoryID = lastId

	ctx.JSON(http.StatusCreated, category)
}

func IndexCategory(ctx *gin.Context) {
	var categories models.Category
	var result []models.Category

	script := `SELECT * FROM categories`
	rows, err := initDb().Query(script)
	checkError(err, "Get Index Failed")
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&categories.CategoryID, &categories.CategoryName, &categories.Details)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, categories)
	}

	ctx.JSON(http.StatusOK, result)
}

func GetCategory(ctx *gin.Context) {
	var category models.Category

	// Convert params id from string to int
	id := ctx.Param("id")
	categoryId, error := strconv.Atoi(id)
	checkError(error, "Convert Failed")

	// Find data with id
	script := fmt.Sprintf(`SELECT * FROM categories WHERE id = %d`, categoryId)
	rows := initDb().QueryRow(script)
	err := rows.Scan(&category.CategoryID, &category.CategoryName, &category.Details)

	// If there isn't match id return 404
	if err != nil && err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})
	} else {
		// If there is match id return data
		ctx.JSON(http.StatusOK, category)
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
