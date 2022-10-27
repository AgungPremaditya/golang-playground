package helpers

import (
	"movies-golang-api/database"
	"strconv"
)

func StrToInt(str string) int {
	int, error := strconv.Atoi(str)
	database.CheckError(error, "Convert Failed")
	return int
}
