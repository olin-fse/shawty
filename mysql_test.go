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
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
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
