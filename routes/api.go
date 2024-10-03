package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lukkas-lukkas/golang-todo-list-api/app/http/controllers"
	"net/http"
)

func ExecuteApi() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", controllers.HelloWorldController)

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		panic(err)
	}
}
