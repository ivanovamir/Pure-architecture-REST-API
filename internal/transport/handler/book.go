package handler

import (
	"encoding/json"
	"fmt"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (h *httpHandler) GetAllBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	booksDTO, err := h.service.GetAllBooks(r.Context())

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

		body, err := json.Marshal(booksDTO)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(pkg.ErrorHandler(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
		return
	}
}

func (h *httpHandler) GetBookByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bookId, err := strconv.Atoi(r.URL.Query().Get("book_id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(pkg.ErrorHandler(fmt.Errorf("invalid book_id")))
		return
	}
	bookDTO, err := h.service.GetBookByID(r.Context(), bookId)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.ErrorHandler(err))
		return
	}

	body, err := json.Marshal(bookDTO)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.ErrorHandler(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return

}
