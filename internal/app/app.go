package app

import (
	"github.com/banggibima/go-chi-restful-api/internal/handlers"
	"github.com/go-chi/chi"
)

type App struct {
	UserHandler *handlers.UserHandler
}

func NewApp(userHandler *handlers.UserHandler) *App {
	return &App{UserHandler: userHandler}
}

func (a *App) SetupRoutes(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", a.UserHandler.GetUsersHandler)
			r.Get("/{id}", a.UserHandler.GetUserByIDHandler)
			r.Post("/", a.UserHandler.CreateUserHandler)
			r.Put("/{id}", a.UserHandler.UpdateUserHandler)
			r.Delete("/{id}", a.UserHandler.DeleteUserHandler)
		})
	})
}
