package handler

import (
	"encoding/json"
	"fmt"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/dto"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"strconv"
)

func (h *httpHandler) GetAllUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	usersDTO, err := h.service.GetAllUsers(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.ErrorHandler(err))
		return
	}
	body, err := json.Marshal(map[string][]*dto.User{"data": usersDTO})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.ErrorHandler(err))
		return
	}

	if len(usersDTO) == 0 {
		w.WriteHeader(http.StatusOK)

		body, err = json.Marshal([]string{})
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

func (h *httpHandler) GetUserByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	userId, err := strconv.Atoi(r.URL.Query().Get("user_id"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.ErrorHandler(fmt.Errorf("%s", errParsTypes)))
		return
	}
	userDTO, err := h.service.GetUserByID(r.Context(), userId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.ErrorHandler(err))
		return
	}

	body, err := json.Marshal(map[string]*dto.User{"data": userDTO})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if userDTO == nil {
		body, err = json.Marshal([]string{})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(pkg.ErrorHandler(err))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return
}

func (h *httpHandler) TakeBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	userId, err := strconv.Atoi(r.Context().Value("userId").(string))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.ErrorHandler(fmt.Errorf("%s", errParsTypes)))
		return
	}

	bookId, err := strconv.Atoi(r.Form.Get("book_id"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.ErrorHandler(fmt.Errorf("%s", errParsTypes)))
		return
	}

	if err := h.service.TakeBook(r.Context(), bookId, userId); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(pkg.ErrorHandler(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func (h *httpHandler) RegisterUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var registerUser *dto.RegisterUser

	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.ErrorHandler(err))
		return
	}

	if err = json.Unmarshal(body, &registerUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(pkg.ErrorHandler(fmt.Errorf("wrong json input")))
		return
	}

	successRegister, err := h.service.RegisterUser(r.Context(), registerUser.Name)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(pkg.ErrorHandler(err))
		return
	}

	response, err := json.Marshal(map[string]dto.SuccessRegister{
		"data": *successRegister,
	})

	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}

func (h *httpHandler) UpdateAccessToken(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var tokensInfoDTO *dto.UpdateTokens

	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.ErrorHandler(err))
		return
	}

	if err = json.Unmarshal(body, &tokensInfoDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(pkg.ErrorHandler(fmt.Errorf("wrong json input")))
		return
	}

	newAccessToken, err := h.service.UpdateAccessToken(r.Context(), tokensInfoDTO.AccessToken, tokensInfoDTO.RefreshToken)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(pkg.ErrorHandler(err))
		return
	}

	response, err := json.Marshal(map[string]map[string]string{
		"data": {
			"access_token": newAccessToken,
		},
	})

	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}

func (h *httpHandler) UpdateRefreshToken(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var tokensInfoDTO *dto.UpdateTokens

	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.ErrorHandler(err))
		return
	}

	if err = json.Unmarshal(body, &tokensInfoDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(pkg.ErrorHandler(fmt.Errorf("wrong json input")))
		return
	}

	newTokens, err := h.service.UpdateRefreshToken(r.Context(), tokensInfoDTO.AccessToken, tokensInfoDTO.RefreshToken)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(pkg.ErrorHandler(err))
		return
	}

	response, err := json.Marshal(newTokens)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}
