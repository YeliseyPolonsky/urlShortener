package link

import (
	"fmt"
	"go-advance/configs"
	"go-advance/pkg/req"
	"go-advance/pkg/res"
	"net/http"
)

//нужен ли конфиг?

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
		dto, err := req.HandleBody[LinkCreateRequest](w, r)
		if err != nil {
			return
		}
		link := NewLink(dto.Url)
		err = h.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) //код неверный?
			return
		}
		res.Json(w, link, 201)
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
		hash := r.PathValue("hash")
		link, err := h.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}
