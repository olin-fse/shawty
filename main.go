package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var s Store

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	connectToDb(&MySqlConfig{
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	})

	r := Handlers()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func connectToDb(config *MySqlConfig) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.username,
		config.password,
		config.host,
		config.port,
		config.database,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// Initialize our store
	s, err = NewStore(db)
	if err != nil {
		panic(err)
	}
}

type Store interface {
	CreateMapping(url, code string, singleUse bool) (bool, error)
	GetUrlForCode(code string) (string, error)
	Close() error
}
