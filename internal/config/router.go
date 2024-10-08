package config

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lukkas-lukkas/golang-todo-list-api/internal/adapters/http/controllers"
)

func CreateRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))

	r.Get("/", controllers.HelloWorldController)
	r.Post("/register", controllers.RegisterController)

	return r
}
