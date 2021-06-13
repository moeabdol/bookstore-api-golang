package utils

import (
	"fmt"

	"github.com/go-playground/validator"
)

// ValidateStruct function to validate structs
func ValidateStruct(object interface{}) []string {
	var valErrors []string
	validate := validator.New()
	err := validate.Struct(object)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			valErrors = append(valErrors, fmt.Sprintf("%s.%s", err.Field(), err.Tag()))
		}
		return valErrors
	}
	return []string{}
}
