package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var s Store

const codeLength = 5

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Connect to database - dsn -> data store name
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s/?parseTime=true",
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

	// Initialize our store
	s, err = NewStore(db)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/generate", generate)
	r.HandleFunc("/{code:[a-zA-Z0-9]{5}}", redirectToUrl)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func generate(w http.ResponseWriter, req *http.Request) {
	requestJson := DecodeJsonForUrl(req)
	generatedCode := RandSeq(codeLength)
	success, _ := s.CreateMapping(requestJson.Url, generatedCode, requestJson.SingleUse)

	if success {
		responseJson := CodeJson{Code: generatedCode}
		json.NewEncoder(w).Encode(responseJson)
	}
}

func redirectToUrl(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	code := vars["code"]

	url, err := s.GetUrlForCode(code)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Could not find this code"))
	} else {
		http.Redirect(w, req, url, 301)
	}
}

type Store interface {
	CreateMapping(url, code string, singleUse bool) (bool, error)
	GetUrlForCode(code string) (string, error)
	Close() error
}
