package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

const (
	requestIDKey = 0
)

type Server struct {
	logger *log.Logger
	router *http.ServeMux
}

func NewServer(
	logger *log.Logger,
) *Server {
	router := http.NewServeMux()

	server := &Server{
		logger: logger,
		router: router,
	}

	server.routes()
	return server
}

func (s *Server) Serve(addr string, wg *sync.WaitGroup) {
	s.logger.Println("Server is starting...")

	srv := &http.Server{
		Addr:         addr,
		Handler:      s.router,
		ErrorLog:     s.logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		s.logger.Println("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			s.logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	wg.Done()
	s.logger.Println("Server is ready to handle requests at", srv.Addr)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Fatalf("Could not listen on %s: %v\n", srv.Addr, err)
	}

	<-done
	s.logger.Println("Server stopped")
}
