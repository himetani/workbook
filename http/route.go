package http

func (s *Server) routes() {
	s.router.HandleFunc("/", s.logging(s.index()))
	s.router.HandleFunc("/auth", s.logging(s.auth()))
}
