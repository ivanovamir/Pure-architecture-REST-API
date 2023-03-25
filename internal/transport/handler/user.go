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

func (h *httpHandler) GetAllUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	usersDTO, err := h.service.GetAllUsers(r.Context())

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

func (h *httpHandler) GetUserByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	userId, err := strconv.Atoi(r.URL.Query().Get("user_id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.ErrorHandler(fmt.Errorf("%s", errParsTypes)))
		return
	}
	userDTO, err := h.service.GetUserByID(r.Context(), userId)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.ErrorHandler(err))
		return
	}

	body, err := json.Marshal(map[string]*dto.User{"data": userDTO})

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if userDTO == nil {
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

func (h *httpHandler) TakeBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	r.ParseForm()
	userId, err := strconv.Atoi(r.Form.Get("user_id"))
	bookId, err := strconv.Atoi(r.Form.Get("book_id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(pkg.ErrorHandler(fmt.Errorf("%s", errParsTypes)))
		return
	}

	if err := h.service.TakeBook(r.Context(), bookId, userId); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(pkg.ErrorHandler(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}
