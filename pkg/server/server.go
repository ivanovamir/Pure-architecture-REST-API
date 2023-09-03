package server

import (
	"context"
	"net"
	"net/http"
)

type server struct {
	srv      *http.Server
	listener *net.Listener
	handler  *http.Handler
}

func NewServer(option ...Option) *server {
	srv := &server{}

	for _, opt := range option {
		opt(srv)
	}

	return srv
}

func (s *server) Run() error {
	if err := s.srv.Serve(*s.listener); err != nil {
		return err
	}
	return nil
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
