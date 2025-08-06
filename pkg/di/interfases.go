package di

import (
	"go-advance/internal/user"
	"net/http"
)

type IStatRepository interface {
	AddClick(linkId uint)
}

type IAuthMiddlware interface {
	IsAuth(next http.Handler) http.Handler
}

type IUserRepository interface {
	Create(user *user.User) error
	FindByEmail(email string) (*user.User, error)
}
