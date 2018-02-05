package main

import (
	"database/sql"
	"fmt"
)

const (
	DRIVER = "url:password@tcp(127.0.0.1:3306)/urlshortener"
	CREATE_MAPPING = "INSERT INTO mappings (original_url, shortened_url, single_use, expired) VALUES(?, ?, ?, ?)"
	FIND_MAPPING = "SELECT (id, original_url, single_use, expired) FROM mappings WHERE shortened_url = ?"
	EXPIRE_MAPPING = "UPDATE mappings SET expired = ? WHERE id = ?"
)

type queryResponse struct {
	Id int
	OriginalUrl string
	SingleUse int
	Expired int
}

var db *sql.DB = nil

func InitDB() *sql.DB {
	db, _ = sql.Open("mysql", DRIVER)

	return db
}

func CreateMapping(url, code string, singleUse bool) bool {
	stmt, err := db.Prepare(CREATE_MAPPING)
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
	err := db.QueryRow(FIND_MAPPING, code).Scan(&res)
	fmt.Println(res)

	if res.SingleUse == 1 {
		if res.Expired == 1 {
			fmt.Println("expired")
			return "", fmt.Errorf("%s has already expired", code)
		} else {
			fmt.Println("will expire")
			stmt, err := db.Prepare(EXPIRE_MAPPING)
			if err != nil {
				fmt.Println(err)
				return "", err
			}

			_, err = stmt.Exec(1, res.Id)
			return res.OriginalUrl, nil
		}
	}

	fmt.Println("no problem")
	return res.OriginalUrl, err
}