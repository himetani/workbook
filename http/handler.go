package http

import (
	"fmt"
	"net/http"
	"sync"
)

func (s *Server) index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		fmt.Fprintf(w, "Hello world\n")
	}
}

func (s *Server) pocketRedirected(wg *sync.WaitGroup) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Printf("Authorized")
		s.client.GetAccessToken()
		fmt.Fprintf(w, "Authorized")
	}
}
