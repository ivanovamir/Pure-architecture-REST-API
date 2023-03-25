package handler

import (
	"fmt"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/service"
	"github.com/julienschmidt/httprouter"
)

// Routes
const (
	getAllBooks      = "books"
	getBookByID      = "book"
	getAllUsers      = "users"
	getUserByID      = "user"
	takeBookByUserID = "take_book"
)

// Errors
const (
	errParsTypes = "error occurred parsing types"
)

type httpHandler struct {
	router  *httprouter.Router
	service *service.Service
}

func NewHttpHandler(router *httprouter.Router, service *service.Service) *httpHandler {
	return &httpHandler{
		router:  router,
		service: service,
	}
}

func (h *httpHandler) Router() {
	h.router.GET(fmt.Sprintf("/%s", getAllBooks), h.GetAllBooks)
	h.router.GET(fmt.Sprintf("/%s", getBookByID), h.GetBookByID)

	h.router.GET(fmt.Sprintf("/%s", getAllUsers), h.GetAllUsers)
	h.router.GET(fmt.Sprintf("/%s", getUserByID), h.GetUserByID)
	h.router.POST(fmt.Sprintf("/%s", takeBookByUserID), h.TakeBook)
}
