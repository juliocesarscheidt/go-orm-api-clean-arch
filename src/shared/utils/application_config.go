package utils

import (
	"fmt"
	"os"
	"strconv"
)

func GetDbConnectionStringConfig() string {
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASS")
	dbDatabase := os.Getenv("MYSQL_DATABASE")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbDatabase)
}

func GetInMemoryDbConfig() bool {
	inMemory, err := strconv.Atoi(os.Getenv("IN_MEMORY_DB"))
	if err != nil {
		return false
	}

	return inMemory == 1
}
