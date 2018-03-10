package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

const codeLength = 5

func generateMapping(w http.ResponseWriter, req *http.Request) {
	requestJson := DecodeJsonForUrl(req)
	generatedCode := RandSeq(codeLength)

	if len(requestJson.Url) == 0 ||
		(requestJson.SingleUse != true && requestJson.SingleUse != false) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

func healthz(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/generate", generateMapping).Methods("POST")
	r.HandleFunc("/{code:[a-zA-Z0-9]{5}}", redirectToUrl).Methods("GET")
	r.HandleFunc("/healthz", healthz).Methods("GET")

	return r
}
