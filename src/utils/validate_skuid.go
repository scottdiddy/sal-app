package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/scottdiddy/sal-app/src/repository"
)

func validateSKUID(f1 validator.FieldLevel) bool {
	_, exists := repository.MerchantData[f1.Field().String()]
	return exists
}
