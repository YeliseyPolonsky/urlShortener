package middlware

import "net/http"

type WriterWrapper struct {
	http.ResponseWriter
	StatusCode int
}

func (wrapper *WriterWrapper) WriteHeader(statusCode int) {
	wrapper.ResponseWriter.WriteHeader(statusCode)
	wrapper.StatusCode = statusCode
}
