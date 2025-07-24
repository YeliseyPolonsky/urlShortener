package hello

import (
	"fmt"
	"io"
	"net/http"
)

type HelloHandler struct{}

func (handler *HelloHandler) Hello() http.HandlerFunc {
	return func(writer http.ResponseWriter, resp *http.Request) {
		io.WriteString(writer, "Hello from Purple School!\n")
		fmt.Println(resp.Method, resp.Header)
	}
}

func NewHelloHandler(router *http.ServeMux) {
	handler := &HelloHandler{}
	router.HandleFunc("/hello", handler.Hello())
}
