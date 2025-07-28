package main

import (
	"fmt"
	"go-advance/configs"
	"go-advance/internal/auth"
	"go-advance/internal/link"
	"go-advance/pkg/db"
	"log"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	db := db.NewDb(config)
	router := http.NewServeMux()

	//Repositories
	linkRepository := link.NewLinkRepository(db)

	//Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: config})
	link.NewLinkHandler(router, link.LinkHandlerDeps{Config: config, LinkRepository: linkRepository})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Start server <- localhost:8080")
	log.Fatal(server.ListenAndServe())
}
