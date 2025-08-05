package di

import "net/http"

type IStatRepository interface {
	AddClick(linkId uint)
}

type IAuthMiddlware interface {
	IsAuth(next http.Handler) http.Handler
}
