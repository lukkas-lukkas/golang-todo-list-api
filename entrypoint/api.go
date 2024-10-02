package entrypoint

import (
	"github.com/go-chi/chi/v5"
	handlers "github.com/lukkas-lukkas/golang-todo-list-api/app/infrastructure/http/handlers"
)

func SetApi(r chi.Router) {
	r.Get("/", handlers.HelloWorld)
}
