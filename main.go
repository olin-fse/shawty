package main

import (
	"html/template"
	"fmt"
	"net/http"
	"path/filepath"
	"encoding/json"
	"database/sql"
	"math/rand"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type UrlJson struct {
	Url string
}

var db *sql.DB = nil
var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	db, _ = sql.Open("mysql",
		"url:password@tcp(127.0.0.1:3306)/urlshortener")
	defer db.Close()

	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./elm/static")))
	r.HandleFunc("/generate", generate)
	r.HandleFunc("/{mapping:[a-zA-Z0-9]{5}}", findMapping)

	http.Handle("/", r)

	fmt.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("elm", "static")
	fp := filepath.Join("templates", "/index.html")

	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func generate(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var requestJson UrlJson
	err := decoder.Decode(&requestJson)
	if err != nil {
		fmt.Println(err)
	}
	defer req.Body.Close()

	// TODO Validate the url they provide...or not
	mapping := randSeq(5)
	stmt, err := db.Prepare("INSERT INTO mappings (original_url, shortened_url) VALUES(?, ?)")
	if (err != nil) {
		fmt.Println(err)
	}

	_, err = stmt.Exec(requestJson.Url, mapping)
	if (err != nil) {
		fmt.Println(err)
	}

	responseJson := UrlJson{Url: "localhost:8080/" + mapping}
	json.NewEncoder(w).Encode(responseJson)
}

func findMapping(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	mapping := vars["mapping"]

	var originalUrl string
	err := db.QueryRow("SELECT original_url FROM mappings WHERE shortened_url = ?", mapping).Scan(&originalUrl)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Could not find this mapping"))
	} else {
		http.Redirect(w, req, originalUrl, 301)
	}
}

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = chars[rand.Intn(len(chars))]
    }
    return string(b)
}
