package main

import (
	"bytes"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"net/http"
	"net/http/httptest"
)

var (
	server      *httptest.Server
	reader      io.Reader
	generateUrl string
	pingUrl     string
)

//func TestCart(t *testing.T) {
//	RegisterFailHandler(Fail)
//	RunSpecs(t, "Test Suite")
//}

func init() {
	server = httptest.NewServer(Handlers())
	fmt.Println(server.URL)
	generateUrl = fmt.Sprintf("%s/generate", server.URL)
	pingUrl = fmt.Sprintf("%s/ping", server.URL)
}

var _ = Describe("Test API Endpoints", func() {
	It("generates a new mapping when POST /generate is hit", func() {
		resp := httptest.NewRecorder()

		request := httptest.NewRequest("POST", generateUrl, nil)

		http.DefaultServeMux.ServeHTTP(resp, request)

		Expect(resp.Result().StatusCode).To(Equal(200))
	})

	FIt("returns 200 on GET /ping", func() {
		fmt.Printf("%s\n", pingUrl)
		resp := httptest.NewRecorder()

		request := httptest.NewRequest("GET", pingUrl, nil)

		http.DefaultServeMux.ServeHTTP(resp, request)

		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		newStr := buf.String()
		fmt.Println("body", newStr)

		Expect(resp.Result().StatusCode).To(Equal(200))
	})
})
