package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"math/rand"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"./service"
	"./utils"
)

type UrlJson struct {
	Url string
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	db := service.InitDB()
	defer db.Close()

	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("./elm/dist")))
	r.Handle("/{jsFile:[a-z]+.js}", http.FileServer(http.Dir("./elm/dist")))
	r.HandleFunc("/generate", generate)
	r.HandleFunc("/{code:[a-zA-Z0-9]{5}}", redirectToUrl)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
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
	generatedCode := utils.RandSeq(5)
	success := service.CreateMapping(requestJson.Url, generatedCode)

	if (success) {
		responseJson := UrlJson{Url: "localhost:8080/" + generatedCode}
		json.NewEncoder(w).Encode(responseJson)
	}
}

func redirectToUrl(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	code := vars["code"]

	url, err := service.GetUrlForCode(code)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Could not find this code"))
	} else {
		http.Redirect(w, req, url, 301)
	}
}

