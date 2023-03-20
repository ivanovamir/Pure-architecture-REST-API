package handler

import (
	"encoding/json"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/dto"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (h *httpHandler) GetAllUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	usersDTO, err := h.service.GetAllUsers(r.Context())

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.ErrorHandler(err))
		return
	} else {
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(pkg.ErrorHandler(err))
			return
		}

		body, err := json.Marshal(map[string][]*dto.User{"data": usersDTO})

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(pkg.ErrorHandler(err))
			return
		}

		if len(usersDTO) == 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			body, err = json.Marshal([]string{})
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(pkg.ErrorHandler(err))
				return
			}
			w.Write(body)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
		return
	}
}
