package main

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var exampleHtml, _ = os.ReadFile("./example_page.html")

type PageTestSuite struct {
	suite.Suite
	Server *httptest.Server
}

func (p *PageTestSuite) TearDownSuite() {
	p.Server.Close()
}

func (p *PageTestSuite) SetupTest() {
	p.Server = newTestServer(exampleHtml)
}

func (p *PageTestSuite) Test() {
	// -- Given
	//

	// -- When
	//

	// -- Then
	//
}

func TestPageTestSuite(t *testing.T) {
	suite.Run(t, new(PageTestSuite))
}

var serverIndexResponse = []byte("hello world\n")

func newTestServer(content []byte) *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write(serverIndexResponse)
	})

	mux.HandleFunc("/html", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write(content)
	})

	return httptest.NewServer(mux)
}
