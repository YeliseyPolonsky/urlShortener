package middlware

import (
	"log"
	"net/http"
	"strings"
)

func IsAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		TOKEN := strings.TrimPrefix(auth, "Bearer ")
		log.Println(TOKEN)
		next.ServeHTTP(w, r)
	})
}
