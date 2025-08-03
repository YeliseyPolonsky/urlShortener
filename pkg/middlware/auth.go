package middlware

import (
	"context"
	"go-advance/configs"
	"go-advance/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	CtxEmailKey key = "CtxEmailKey "
)

func IsAuth(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		TOKEN := strings.TrimPrefix(auth, "Bearer ")
		_, data := jwt.NewJWT(config.Auth.Secret).Parse(TOKEN)
		ctx := context.WithValue(r.Context(), CtxEmailKey, data.Email)
		newReq := r.WithContext(ctx)
		next.ServeHTTP(w, newReq)
	})
}
