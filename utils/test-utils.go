package utils

import (
	"net/http"
	"net/http/httptest"
)

func Setup(handler func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}))
}

func Teardown(s *httptest.Server) {
	s.Close()
}
