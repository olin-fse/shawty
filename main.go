package main

import (
	"net/http"
	"encoding/json"
	"math/rand"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const codeLength = 5

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	db := InitDB()
	defer db.Close()

	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("./frontend/public")))
	r.Handle("/{jsFile:[a-z]+.js}", http.FileServer(http.Dir("./frontend/public")))
	r.HandleFunc("/generate", generate)
	r.HandleFunc("/{code:[a-zA-Z0-9]{5}}", redirectToUrl)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func generate(w http.ResponseWriter, req *http.Request) {
	url := DecodeJsonForUrl(req)
	generatedCode := RandSeq(codeLength)
	success := CreateMapping(url, generatedCode)

	if success {
		responseJson := UrlJson{Url: "localhost:8080/" + generatedCode}
		json.NewEncoder(w).Encode(responseJson)
	}
}

func redirectToUrl(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	code := vars["code"]

	url, err := GetUrlForCode(code)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Could not find this code"))
	} else {
		http.Redirect(w, req, url, 301)
	}
}

