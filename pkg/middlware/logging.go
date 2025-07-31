package middlware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		writerWrapper := &WriterWrapper{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(writerWrapper, r)
		log.Println(r.Method, r.URL, time.Since(start), writerWrapper.StatusCode)
	})
}
