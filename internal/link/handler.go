package link

import (
	"go-advance/configs"
	"go-advance/pkg/middlware"
	"go-advance/pkg/req"
	"go-advance/pkg/res"
	"net/http"
	"strconv"

	"gorm.io/gorm"
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
	router.Handle("PATCH /link/{id}", middlware.IsAuth(h.Update(), deps.Config))
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
		// for h.LinkRepository.IsExist("hash", link.Hash) {
		// 	link.GenerateHash()
		// }
		for {
			existedLink, _ := h.LinkRepository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}
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
		// email := r.Context().Value(middlware.CtxEmailKey)
		// log.Println(email)
		dto, err := req.HandleBody[LinkUpdateRequest](w, r)
		if err != nil {
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
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}
