package main

import (
	"fmt"
	"go-advance/configs"
	"go-advance/internal/auth"
	"go-advance/internal/link"
	"log"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	fmt.Printf("Start server")
	router := http.NewServeMux()

	//Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: config})
	link.NewLinkHandler(router, link.LinkHandlerDeps{Config: config})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
