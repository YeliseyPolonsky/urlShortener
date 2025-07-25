package auth

import (
	"fmt"
	"go-advance/configs"
	"go-advance/pkg/req"
	"go-advance/pkg/res"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
}

type AuthHandlerDeps struct {
	*configs.Config
}

func NewAuthHandler(r *http.ServeMux, deps AuthHandlerDeps) {
	h := &AuthHandler{
		Config: deps.Config,
	}
	r.HandleFunc("POST /auth/login", h.Login())
	r.HandleFunc("POST /auth/register", h.Register())
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dto, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(*dto)
		data := LoginResponse{
			Token: h.Config.Auth.Secret,
		}

		res.Json(w, data, 201)
	}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dto, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(*dto)
		data := RegisterResponse{
			Token: fmt.Sprint("Register:", h.Config.Auth.Secret),
		}

		res.Json(w, data, 201)
	}
}
