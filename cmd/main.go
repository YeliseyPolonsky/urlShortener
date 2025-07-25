package main

import (
	"fmt"
	"go-advance/configs"
	"go-advance/internal/auth"
	"log"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	fmt.Printf("Start server")
	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: config})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
