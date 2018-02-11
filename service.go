package main

import (
	"database/sql"
	"fmt"
)

const (
	DRIVER        = "url:password@tcp(127.0.0.1:3306)/urlshortener"
	CREATEMAPPING = "INSERT INTO mappings (original_url, shortened_url, single_use, expired) VALUES(?, ?, ?, ?)"
	FINDMAPPING   = "SELECT id, original_url, single_use, expired FROM mappings WHERE shortened_url = ?"
	EXPIREMAPPING = "UPDATE mappings SET expired = ? WHERE id = ?"
)

type queryResponse struct {
	id          int
	originalUrl string
	singleUse   int
	expired     int
}

var db *sql.DB = nil

func InitDB() *sql.DB {
	db, _ = sql.Open("mysql", DRIVER)

	return db
}

func CreateMapping(url, code string, singleUse bool) bool {
	stmt, err := db.Prepare(CREATEMAPPING)
	if err != nil {
		fmt.Println(err)
		return false
	}

	_, err = stmt.Exec(url, code, singleUse, false)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func GetUrlForCode(code string) (string, error) {
	var res queryResponse
	err := db.QueryRow(FINDMAPPING, code).Scan(&res.id, &res.originalUrl, &res.singleUse, &res.expired)
	if err != nil {
		return "", fmt.Errorf("%s was not found", code)
	}

	if res.singleUse == 1 {
		if res.expired == 1 {
			return "", fmt.Errorf("%s has already expired", code)
		} else {
			stmt, err := db.Prepare(EXPIREMAPPING)
			if err != nil {
				fmt.Println(err)
				return "", err
			}

			_, err = stmt.Exec(1, res.id)
			return res.originalUrl, nil
		}
	}

	return res.originalUrl, nil
}