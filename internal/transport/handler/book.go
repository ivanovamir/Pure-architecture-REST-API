package handler

import (
	"encoding/json"
	"fmt"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/dto"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (h *httpHandler) GetAllBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	booksDTO, err := h.service.GetAllBooks(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.ErrorHandler(err))
		return
	}
	body, err := json.Marshal(map[string][]*dto.Book{"data": booksDTO})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.ErrorHandler(err))
		return
	}

	if len(booksDTO) == 0 {
		w.WriteHeader(http.StatusOK)
		body, err = json.Marshal(map[string][]string{"data": []string{}})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(pkg.ErrorHandler(err))
			return
		}
		w.Write(body)
		return
	}

	w.Write(body)
	return

}

func (h *httpHandler) GetBookByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	bookId, err := strconv.Atoi(r.URL.Query().Get("book_id"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.ErrorHandler(fmt.Errorf("%s", errParsTypes)))
		return
	}
	bookDTO, err := h.service.GetBookByID(r.Context(), bookId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.ErrorHandler(err))
		return
	}

	body, err := json.Marshal(map[string]*dto.Book{"data": bookDTO})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.ErrorHandler(err))
		return
	}

	if bookDTO.Id == "" {
		w.WriteHeader(http.StatusOK)

		body, err = json.Marshal(map[string]struct{}{"data": struct{}{}})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(pkg.ErrorHandler(err))
			return
		}
		w.Write(body)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return
}
