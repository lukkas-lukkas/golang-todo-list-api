package controllers

import (
	"fmt"
	"github.com/lukkas-lukkas/golang-todo-list-api/internal/adapters/http/requests"
	"net/http"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {

	var request requests.RegisterRequest

	if validated := requests.Validated(request, w); !validated {
		return
	}

	fmt.Println(request)

	w.Write([]byte("User created"))
}
