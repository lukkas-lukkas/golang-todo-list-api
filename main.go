package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lukkas-lukkas/golang-todo-list-api/entrypoint"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	entrypoint.SetApi(r)

	http.ListenAndServe(":8000", r)
}
