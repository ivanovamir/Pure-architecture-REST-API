package middleware

import (
	"context"
	"fmt"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/service"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

type AppHandler func(w http.ResponseWriter, r *http.Request, params httprouter.Params)

type middleware struct {
	service *service.Service
}

type Middleware interface {
	UserIdentity(h AppHandler) httprouter.Handle
}

func NewMiddleware(service *service.Service) Middleware {
	return &middleware{service: service}
}

func (m *middleware) UserIdentity(h AppHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(pkg.ErrorHandler(fmt.Errorf("empty auth user")))
			return
		}

		headerParts := strings.Split(authHeader, " ")

		if headerParts[0] != "Bearer" || len(headerParts[1]) == 0 || len(headerParts) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(pkg.ErrorHandler(fmt.Errorf("invalid auth user")))
			return
		}

		userId, err := m.service.UserValidate(headerParts[1])

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(pkg.ErrorHandler(err))
			return
		}

		h(w, r.WithContext(context.WithValue(r.Context(), "userId", userId)), params)
	}
}
