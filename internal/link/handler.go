package link

import (
	"go-advance/configs"
	"net/http"
)

type LinkHandler struct {
	*configs.Config
}

type LinkHandlerDeps struct {
	*configs.Config
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	h := &LinkHandler{
		deps.Config,
	}

	router.HandleFunc("POST /link", h.Create())
	router.HandleFunc("PATCH /link/{id}", h.Update())
	router.HandleFunc("DELETE /link/{id}", h.Delete())
	router.HandleFunc("GET /{alias}", h.GoTo())
}

func (h *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (h *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (h *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (h *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
