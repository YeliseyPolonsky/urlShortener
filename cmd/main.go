package main

import (
	"fmt"
	"go-advance/configs"
	"go-advance/internal/auth"
	"log"
	"net/http"
)

func main() {
	_ = configs.LoadConfig()
	fmt.Printf("Start server")
	router := http.NewServeMux()

	auth.NewAuthHandler(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
