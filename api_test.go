package main

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"net/http"
	"net/http/httptest"
	//"strings"
	//"log"
	"strings"
	"log"
)

var (
	server      *httptest.Server
	reader      io.Reader
	generateUrl string
	healthUrl   string
)

func init() {
	server = httptest.NewServer(Handlers())
	generateUrl = fmt.Sprintf("%s/generate", server.URL)
	healthUrl = fmt.Sprintf("%s/healthz", server.URL)
}

var _ = Describe("Test API Endpoints", func() {
	It("generates a new mapping when POST /generate is hit", func() {
		json := `{"url": "https://google.com", "singleUse": false}`
		reader = strings.NewReader(json) //Convert string to reader

		request, err := http.NewRequest("POST", generateUrl, reader)
		if err != nil {
			log.Fatal(err)
		}

		res, err := http.DefaultClient.Do(request)
		if err != nil {
			log.Fatal(err)
		}

		Expect(res.StatusCode).To(Equal(200))
	})

	It("returns 200 on GET /ping", func() {
		request, _ := http.NewRequest("GET", healthUrl, nil)

		res, _ := http.DefaultClient.Do(request)

		Expect(res.StatusCode).To(Equal(200))
	})
})
