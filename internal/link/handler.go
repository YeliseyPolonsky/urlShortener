package link

import (
	"go-advance/configs"
	"go-advance/pkg/di"
	"go-advance/pkg/event"
	"go-advance/pkg/req"
	"go-advance/pkg/res"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type LinkHandler struct {
	*LinkRepository
	*event.EventBus
	di.IAuthMiddlware
}

type LinkHandlerDeps struct {
	*configs.Config
	*LinkRepository
	*event.EventBus
	di.IAuthMiddlware
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	h := &LinkHandler{
		deps.LinkRepository,
		deps.EventBus,
		deps.IAuthMiddlware,
	}

	router.HandleFunc("POST /link", h.Create())
	router.Handle("PATCH /link/{id}", h.Update())
	router.HandleFunc("DELETE /link/{id}", h.Delete())
	router.HandleFunc("GET /{hash}", h.GoTo())
	router.Handle("GET /link", h.GetAll())
}

func (h *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dto, err := req.HandleBody[LinkCreateRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		link := NewLink(dto.Url)
		for {
			existedLink, _ := h.LinkRepository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}
		err = h.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, link, 201)
	}
}

func (h *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// email := r.Context().Value(middlware.CtxEmailKey)
		// log.Println(email)
		dto, err := req.HandleBody[LinkUpdateRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		link := &Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   dto.Url,
			Hash:  dto.Hash,
		}
		err = h.LinkRepository.Update(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, link, 200)
	}
}

func (h *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = h.LinkRepository.GetByID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = h.LinkRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, "delete success", 200)
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
		go h.EventBus.Publish(event.Event{
			Type: event.LinkVisited,
			Data: link.ID,
		})
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}

func (h *LinkHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			http.Error(w, "Invalid limit", http.StatusBadRequest)
			return
		}

		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			http.Error(w, "Invalid offset", http.StatusBadRequest)
			return
		}
		links := h.LinkRepository.GetAll(limit, offset)
		count := h.LinkRepository.Count()
		res.Json(w, GetAllLinksResponse{
			Links: links,
			Count: count,
		}, http.StatusOK)
	}
}
