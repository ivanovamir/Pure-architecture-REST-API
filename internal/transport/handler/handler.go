package handler

import (
	"fmt"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/service"
	"github.com/julienschmidt/httprouter"
)

const (
	getAllBooks = "books"
	getBookByID = "book"
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
}
