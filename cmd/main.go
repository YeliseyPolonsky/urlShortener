package main

import (
	"fmt"
	"go-advance/configs"
	"go-advance/internal/hello"
	"log"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	fmt.Printf("Start server")
	router := http.NewServeMux()
	hello.NewHelloHandler(router)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
