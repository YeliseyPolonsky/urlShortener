package main

import (
	"fmt"
	"go-advance/configs"
	"go-advance/internal/auth"
	"go-advance/internal/link"
	"go-advance/internal/stat"
	"go-advance/internal/user"
	"go-advance/pkg/db"
	"go-advance/pkg/event"
	"go-advance/pkg/jwt"
	"go-advance/pkg/middlware"
	"log"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	db := db.NewDb(config)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	//Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	//Services
	authService := auth.NewUserService(userRepository)
	jwtService := jwt.NewJWT(config.Auth.Secret)
	statService := stat.NewStatService(stat.StatServiceDep{
		StatRepository: statRepository,
		EventBus:       eventBus,
	})
	go statService.AddClick()

	var authURLs []string
	//Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      config,
		AuthService: authService,
	})
	authURLs = link.NewLinkHandler(router, link.LinkHandlerDeps{
		Config:         config,
		LinkRepository: linkRepository,
		EventBus:       eventBus,
	})
	stat.NewStatHandler(router, stat.StatHandlerDep{
		StatRepository: statRepository,
		Config:         config,
	})

	//Stateful middlware
	authMiddlware := middlware.NewAuthMiddleware(middlware.AuthMiddlewareDeps{
		JWTService:  jwtService,
		AuthPaterns: authURLs,
	})

	//Middlwares
	stack := middlware.Chain(
		middlware.CORS,
		middlware.Logging,
		authMiddlware.IsAuth,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	fmt.Println("Start server <- localhost:8080")
	log.Fatal(server.ListenAndServe())
}
