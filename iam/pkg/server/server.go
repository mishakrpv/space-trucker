package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type Server struct {
	stopChan chan bool

	HTTPServer *http.Server
}

// NewServer returns an initialized Server.
func NewServer(httpServer *http.Server) *Server {
	return &Server{
		stopChan:   make(chan bool, 1),
		HTTPServer: httpServer,
	}
}

// Start starts the server and Stop/Close it when context is Done.
func (s *Server) Start(ctx context.Context) {
	go func() {
		<-ctx.Done()
		logger := log.Ctx(ctx)
		logger.Info().Msg("I have to go...")
		logger.Info().Msg("Stopping server gracefully")
		s.Stop()
	}()

	log.Info().Msgf("Server is listening on: %s", s.HTTPServer.Addr)

	if err := s.HTTPServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Error().Err(err).Msg("Something's wrong, I can feel it")
	}
}

// Wait blocks until the server shutdown.
func (s *Server) Wait() {
	<-s.stopChan
}

// Stop stops the server.
func (s *Server) Stop() {
	defer log.Info().Msg("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.HTTPServer.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Shutdown returned an error")
	}

	s.stopChan <- true
}

// Close destroys the server.
func (s *Server) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	go func(ctx context.Context) {
		<-ctx.Done()
		if errors.Is(ctx.Err(), context.Canceled) {
			return
		} else if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			panic("Timeout while stopping, killing instance")
		}
	}(ctx)

	close(s.stopChan)

	cancel()
}
