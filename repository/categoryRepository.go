package repository

import (
	"database/sql"
	"fmt"
	"log"
	"movies-golang-api/database"
	"movies-golang-api/models"
)

func CreateCategoryQuery(category models.Category) models.Category {
	var result models.Category

	script := `INSERT INTO categories (category_name, details) VALUES (?, ?)`

	rows, err := database.InitDb().Exec(script, category.CategoryName, category.Details)
	database.CheckError(err, "Insert Failed")
	lastId, _ := rows.LastInsertId()

	result.CategoryID = lastId
	result.CategoryName = category.CategoryName
	result.Details = category.Details

	return result
}

func IndexCategoryQuery() []models.Category {
	var data models.Category
	var result []models.Category

	script := `SELECT * FROM categories`
	rows, err := database.InitDb().Query(script)
	database.CheckError(err, "Get Index Failed")
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&data.CategoryID, &data.CategoryName, &data.Details)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, data)
	}

	return result
}

func GetCategoryQuery(id int) models.Category {
	var result models.Category

	script := fmt.Sprintf(`SELECT * FROM categories WHERE id = %d`, id)
	rows := database.InitDb().QueryRow(script)
	err := rows.Scan(&result.CategoryID, &result.CategoryName, &result.Details)

	if err != nil && err == sql.ErrNoRows {
		return result
	} else {
		database.CheckError(err, "Get Data Failed")
	}

	return result
}
