package database

import (
	"database/sql"
	"movies-golang-api/helpers"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitDb() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/movie_api")
	helpers.CheckError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
