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

func writeUnauthStatus(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuth(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			writeUnauthStatus(w)
			return
		}
		TOKEN := strings.TrimPrefix(auth, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(TOKEN)
		if !isValid {
			writeUnauthStatus(w)
			return
		}
		ctx := context.WithValue(r.Context(), CtxEmailKey, data.Email)
		newReq := r.WithContext(ctx)
		next.ServeHTTP(w, newReq)
	})
}
