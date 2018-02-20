package main

import (
	"database/sql"
	"fmt"
	"os"
)

func newTestStore() *StoreStruct {
	// Connect to the database
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/?parseTime=true",
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_HOST"),
		os.Getenv("TEST_DB_PORT"),
		os.Getenv("TEST_DB_NAME"),
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	s, err := NewStore(db)
	if err != nil {
		panic(err)
	}
	return s
}
