package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/banggibima/go-chi-restful-api/internal/app"
	"github.com/banggibima/go-chi-restful-api/internal/config"
	"github.com/banggibima/go-chi-restful-api/internal/database"
	"github.com/banggibima/go-chi-restful-api/internal/handlers"
	"github.com/banggibima/go-chi-restful-api/internal/repositories"
	"github.com/banggibima/go-chi-restful-api/internal/usecases"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading configuration: %v", err)
	}

	db, err := database.NewDBConnection()
	if err != nil {
		log.Fatalf("error establishing database connection: %v", err)
	}

	userRepo := repositories.NewUserRepository(db)
	userUseCase := usecases.NewUserUseCase(userRepo)
	userHandler := handlers.NewUserHandler(userUseCase)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	myApp := app.NewApp(userHandler)
	myApp.SetupRoutes(r)

	port := fmt.Sprintf(":%d", cfg.Server.Port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("error starting the application: %v", err)
	}
}
