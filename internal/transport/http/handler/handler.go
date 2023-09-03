package handler

import (
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type HttpHandler struct {
	router  *httprouter.Router
	service *service.Service
}

func NewHttpHandler(router *httprouter.Router, service *service.Service) *HttpHandler {
	return &HttpHandler{
		router:  router,
		service: service,
	}
}

func (h *HttpHandler) Router() {
	h.router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("hello world!"))
	})
}
