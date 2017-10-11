package service

import (
	"database/sql"
	"fmt"
)

const (
	DRIVER = "url:password@tcp(127.0.0.1:3306)/urlshortener"
	CREATE_MAPPING = "INSERT INTO mappings (original_url, shortened_url) VALUES(?, ?)"
	FIND_MAPPING = "SELECT original_url FROM mappings WHERE shortened_url = ?"
)

var db *sql.DB = nil

func InitDB() *sql.DB {
	db, _ = sql.Open("mysql", DRIVER)

	return db
}

func CreateMapping(url, code string) bool {
	stmt, err := db.Prepare(CREATE_MAPPING)
	if (err != nil) {
		fmt.Println(err)
		return false
	}

	_, err = stmt.Exec(url, code)
	if (err != nil) {
		fmt.Println(err)
		return false
	}

	return true
}

func GetUrlForCode(code string) (string, error) {
	var url string
	err := db.QueryRow(FIND_MAPPING, code).Scan(&url)

	return url, err
}