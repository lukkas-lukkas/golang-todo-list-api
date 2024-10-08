package controllers

import (
	"encoding/json"
	"github.com/lukkas-lukkas/golang-todo-list-api/internal/adapters/db"
	encrypt2 "github.com/lukkas-lukkas/golang-todo-list-api/internal/adapters/encrypt"
	"github.com/lukkas-lukkas/golang-todo-list-api/internal/adapters/http/requests"
	"github.com/lukkas-lukkas/golang-todo-list-api/internal/adapters/http/responses"
	"github.com/lukkas-lukkas/golang-todo-list-api/internal/usecases/register_user"
	"log"
	"net/http"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {

	var request requests.RegisterRequest

	json.NewDecoder(r.Body).Decode(&request)

	if validated := requests.Validated(request, w); !validated {
		return
	}

	repo := db.NewUserRepository(db.DB)
	encrypt := encrypt2.NewBcryptService()
	useCase := register_user.NewRegisterUserService(repo, encrypt)

	user, err := useCase.CreateUser(register_user.UserDto{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		log.Println("Error registering user", err)
		http.Error(w, "Internal server error", http.StatusBadGateway)
		return
	}

	json.NewEncoder(w).Encode(responses.RegisterUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}
