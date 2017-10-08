package main

import (
	"html/template"
	"fmt"
	"net/http"
	"path/filepath"
	"encoding/json"
)

type requestBody struct {
	UserUrl string
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./elm/static")))
	http.HandleFunc("/generate", generate)

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
	var body requestBody
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	fmt.Println(body.UserUrl)
}
