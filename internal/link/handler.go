package link

import (
	"fmt"
	"go-advance/configs"
	"net/http"
)

type LinkHandler struct {
	*configs.Config
	*LinkRepository
}

type LinkHandlerDeps struct {
	*configs.Config
	*LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	h := &LinkHandler{
		deps.Config,
		deps.LinkRepository,
	}

	router.HandleFunc("POST /link", h.Create())
	router.HandleFunc("PATCH /link/{id}", h.Update())
	router.HandleFunc("DELETE /link/{id}", h.Delete())
	router.HandleFunc("GET /{hash}", h.GoTo())
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
		id := r.PathValue("id")
		fmt.Println(id)
	}
}

func (h *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
