package auth

import (
	"go-advance/configs"
	"go-advance/pkg/jwt"
	"go-advance/pkg/req"
	"go-advance/pkg/res"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
	*AuthService
}

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(r *http.ServeMux, deps AuthHandlerDeps) {
	h := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	r.HandleFunc("POST /auth/login", h.Login())
	r.HandleFunc("POST /auth/register", h.Register())
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dto, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}
		err = h.AuthService.Login(dto.Email, dto.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		token, err := jwt.NewJWT(h.Config.Auth.Secret).Create(dto.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response := LoginResponse{
			Token: token,
		}

		res.Json(w, response, 201)
	}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dto, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			return
		}

		err = h.AuthService.Register(dto.Email, dto.Password, dto.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := jwt.NewJWT(h.Config.Auth.Secret).Create(dto.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response := RegisterResponse{
			Token: token,
		}

		res.Json(w, response, 201)
	}
}
