package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()
	var request UserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	errors := validate.Struct(request)
	if errors != nil {
		var errorResponse []ValidationErrorResponse
		for _, err := range errors.(validator.ValidationErrors) {
			message := fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", err.Field(), err.Tag())
			errorResponse = append(errorResponse, ValidationErrorResponse{
				err.Field(),
				message,
			})
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest) // 400 Bad Request
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": errorResponse,
		})
		return
	}

	fmt.Println(request)

	w.Write([]byte("User created"))
}

type UserRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
