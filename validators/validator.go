package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Field string
	Tag   string
	Value string
}

func ValidateStruct(any any) []*ErrorResponse {
	var validate = validator.New()
	var errors []*ErrorResponse
	err := validate.Struct(any)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			element := ErrorResponse{
				Field: strings.ToLower(err.Field()),
				Tag:   err.Tag(),
				Value: err.Param(),
			}
			errors = append(errors, &element)
		}
	}
	return errors
}
