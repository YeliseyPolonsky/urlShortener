package auth

import (
	"fmt"
	"io"
	"net/http"
)

type AuthHandler struct{}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Login LOL\n")
		fmt.Println(r.Method, "Login")
	}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Registration LOL\n")
		fmt.Println(r.Method, "Registration")
	}
}

func NewAuthHandler(r *http.ServeMux) {
	h := &AuthHandler{}
	r.HandleFunc("POST /auth/login", h.Login())
	r.HandleFunc("POST /auth/register", h.Register())
}
