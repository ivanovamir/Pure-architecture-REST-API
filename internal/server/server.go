package server

import (
	"net"
	"net/http"
)

type server struct {
	srv      *http.Server
	listener net.Listener
}

func NewServer(srv *http.Server, listener net.Listener) *server {
	return &server{srv: srv, listener: listener}
}

func (s *server) Run() error {
	if err := s.srv.Serve(s.listener); err != nil {
		return err
	}
	return nil
}
