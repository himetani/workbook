package http

import "net/http"

func (s *Server) logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
		next.ServeHTTP(w, r)
	}
}
