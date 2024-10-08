package requests

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func Validated(request interface{}, w http.ResponseWriter) bool {
	validate := validator.New()

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
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": errorResponse,
		})
		return false
	}
	return true
}
