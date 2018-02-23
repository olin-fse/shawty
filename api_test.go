package main

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"log"
	"os"
	"encoding/json"
)

var (
	server      *httptest.Server
	reader      io.Reader
	generateUrl string
	healthUrl   string
)

type JsonCodeResponse struct {
	Code string `json:"code"`
}

func init() {
	connectToDb(&MySqlConfig{
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_HOST"),
		os.Getenv("TEST_DB_PORT"),
		os.Getenv("TEST_DB_NAME"),
	})
	server = httptest.NewServer(Handlers())
	generateUrl = fmt.Sprintf("%s/generate", server.URL)
	healthUrl = fmt.Sprintf("%s/healthz", server.URL)
}

func makeGenerateRequest(url string, singleUse bool) *http.Response {
	reqJson := fmt.Sprintf(`{"url": "%s", "singleUse": %t}`, url, singleUse)
	reader = strings.NewReader(reqJson)

	req, err := http.NewRequest("POST", generateUrl, reader)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func makeCodeRequest(code string) *http.Response {
	url := fmt.Sprintf("%s/%s", server.URL, code)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

var _ = Describe("Test API Endpoints", func() {
	It("generates a new mapping when POST /generate is hit", func() {
		res := makeGenerateRequest("https://google.com", false)
		Expect(res.StatusCode).To(Equal(200))
	})

	It("returns 400 if no url is provided", func() {
		res := makeGenerateRequest("", false)
		Expect(res.StatusCode).To(Equal(400))
	})

	It("singleUse flow", func() {
		res := makeGenerateRequest("https://google.com", true)
		Expect(res.StatusCode).To(Equal(200))

		var resJson JsonCodeResponse
		defer res.Body.Close()
		json.NewDecoder(res.Body).Decode(&resJson)

		// First use
		makeCodeRequest(resJson.Code)

		// Second use
		res = makeCodeRequest(resJson.Code)
		Expect(res.StatusCode).To(Equal(http.StatusNotFound))
	})
})
