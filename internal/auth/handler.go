package auth

import (
	"fmt"
	"go-advance/configs"
	"go-advance/pkg/res"
	"io"
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
		fmt.Println(r.Method, "Login")
		data := LoginResponse{
			Token: h.Config.Auth.Secret,
		}

		res.Json(w, data, 201)
	}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Registration LOL\n")
		fmt.Println(r.Method, "Registration")
	}
}
