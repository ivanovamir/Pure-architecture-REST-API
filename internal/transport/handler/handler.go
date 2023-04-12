package handler

import (
	"fmt"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/service"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/transport/handler/middleware"
	"github.com/julienschmidt/httprouter"
)

// Routes
const (
	getAllBooks      = "books"
	getBookByID      = "book"
	getAllUsers      = "users"
	getUserByID      = "user"
	takeBookByUserID = "take_book"
	registerUser     = "sign-in"
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
	h.router.GET(fmt.Sprintf("/%s", getAllBooks), middleware.UserIdentity(h.GetAllBooks))
	h.router.GET(fmt.Sprintf("/%s", getBookByID), middleware.UserIdentity(h.GetBookByID))
	h.router.GET(fmt.Sprintf("/%s", getAllUsers), middleware.UserIdentity(h.GetAllUsers))
	h.router.GET(fmt.Sprintf("/%s", getUserByID), middleware.UserIdentity(h.GetUserByID))
	h.router.POST(fmt.Sprintf("/%s", takeBookByUserID), middleware.UserIdentity(h.TakeBook))

	h.router.POST(fmt.Sprintf("/%s", registerUser), h.RegisterUser)
}
