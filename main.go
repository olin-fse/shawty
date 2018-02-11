package main

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const codeLength = 5

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	db := InitDB()
	defer db.Close()

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
	success := CreateMapping(requestJson.Url, generatedCode, requestJson.SingleUse)

	if success {
		responseJson := CodeJson{Code: generatedCode}
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
