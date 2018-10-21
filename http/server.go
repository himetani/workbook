package http

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/himetani/workbook/pocket"
)

const (
	requestIDKey = 0
)

type Server struct {
	client *pocket.Client
	logger *log.Logger
	router *http.ServeMux
}

func NewServer(
	client *pocket.Client,
	logger *log.Logger,
) *Server {
	router := http.NewServeMux()

	server := &Server{
		client: client,
		logger: logger,
		router: router,
	}

	return server
}

func (s *Server) Serve(addr string, svrStartUp, authCode *sync.WaitGroup, ctx context.Context) {
	s.logger.Println("Server is starting...")
	s.routes(authCode)

	srv := &http.Server{
		Addr:         addr,
		Handler:      s.router,
		ErrorLog:     s.logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan bool)

	go func() {
		<-ctx.Done()
		s.logger.Println("Server is shutting down...")

		innerCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(innerCtx); err != nil {
			s.logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	svrStartUp.Done()
	s.logger.Println("Server is ready to handle requests at", srv.Addr)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Fatalf("Could not listen on %s: %v\n", srv.Addr, err)
	}

	<-done
	s.logger.Println("Server stopped")
}
