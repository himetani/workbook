package http

import (
	"fmt"
	"net/http"
)

func (s *Server) index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		fmt.Fprintf(w, "Hello world\n")
	}
}

func (s *Server) auth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Printf("Redirect to /")
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}
