package http

import "sync"

func (s *Server) routes(wg *sync.WaitGroup) {
	s.router.HandleFunc("/", s.logging(s.index()))
	s.router.HandleFunc("/pocket/redirected", s.logging(s.pocketRedirected(wg)))
}
