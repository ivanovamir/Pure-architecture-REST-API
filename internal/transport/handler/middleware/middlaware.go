package middleware

import (
	"fmt"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

type appHandler func(w http.ResponseWriter, r *http.Request, params httprouter.Params)

func UserIdentity(h appHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			w.Write(pkg.ErrorHandler(fmt.Errorf("empty auth user")))
			return
		}

		headerParts := strings.Split(authHeader, " ")

		if headerParts[0] != "Bearer" || len(headerParts[1]) == 0 || len(headerParts) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			w.Write(pkg.ErrorHandler(fmt.Errorf("invalid auth user")))
			return
		}

		userId, err :=

		if header != "" {
			headerParts := strings.Split(header, " ")

			if len(headerParts) != 2 {
				http.Error(w, "empty auth user", http.StatusUnauthorized)
				return
			}

			userId, err := h.service.ParseToken(headerParts[1])
			if err != nil {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			h(w, r, params)
		} else {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}
	}
}
