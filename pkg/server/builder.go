package server

import (
	"net"
	"net/http"
)

type Option func(*server)

func WithListener(ln *net.Listener) Option {
	return func(s *server) {
		s.listener = ln
	}
}

func WithSrv(srv *http.Server) Option {
	return func(s *server) {
		s.srv = srv
	}
}
