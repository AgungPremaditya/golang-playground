package controllers

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func initDb() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/movie_api")

	checkError(err, "sql.Open failed")

	return db
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
