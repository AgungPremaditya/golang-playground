package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDb() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/movie_api")

	CheckError(err, "sql.Open failed")

	return db
}

func CheckError(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
