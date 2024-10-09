package controllers

import (
	"encoding/json"
	"github.com/lukkas-lukkas/golang-todo-list-api/app/application/register_user"
	"github.com/lukkas-lukkas/golang-todo-list-api/app/infrastructure/helpers"
	"github.com/lukkas-lukkas/golang-todo-list-api/app/infrastructure/http/api/requests"
	"github.com/lukkas-lukkas/golang-todo-list-api/app/infrastructure/http/api/requests/validators"
	"github.com/lukkas-lukkas/golang-todo-list-api/app/infrastructure/http/api/responses"
	"github.com/lukkas-lukkas/golang-todo-list-api/app/infrastructure/persistence/mysql"
	"log"
	"net/http"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {

	var request requests.RegisterRequest

	json.NewDecoder(r.Body).Decode(&request)

	if validated := validators.Validated(request, w); !validated {
		return
	}

	repo := mysql.NewUserRepository(mysql.Connection)
	encrypt := helpers.NewBcryptEncryptor()
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
