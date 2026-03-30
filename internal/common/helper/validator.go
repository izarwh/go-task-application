package helper

import "github.com/go-playground/validator/v10"

func ValidateStruct(validators *validator.Validate, s interface{}) error {
	return validators.Struct(s)
}
