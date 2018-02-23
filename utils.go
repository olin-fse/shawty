package main

import (
	"math/rand"
	"net/http"
	"encoding/json"
	"fmt"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type CreateCodeRequestJson struct {
	Url string `json:"url"`
	SingleUse bool `json:"singleUse"`
}

type CodeJson struct {
	Code string `json:"code"`
}

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func DecodeJsonForUrl(req *http.Request) CreateCodeRequestJson {
	decoder := json.NewDecoder(req.Body)
	var requestJson CreateCodeRequestJson
	err := decoder.Decode(&requestJson)
	if err != nil {
		fmt.Println(err)
	}
	defer req.Body.Close()
	return requestJson
}