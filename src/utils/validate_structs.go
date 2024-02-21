package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateStructBody(reqBody interface{}) error {
	validate := validator.New()
	validate.RegisterValidation("validSKUID", validateSKUID)
	err := validate.Struct(reqBody)
	if err != nil {
		validationError, ok := err.(validator.ValidationErrors)
		if ok {
			for _, err := range validationError {
				msg := "Field " + err.Field() + " failed on the " + err.Tag() + " tag."
				return errors.New(msg)
			}
		} else {
			return err
		}
	}
	return nil
}
