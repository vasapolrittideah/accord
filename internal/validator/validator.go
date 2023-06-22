package validate

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/vasapolrittideah/accord/internal/response"
)

var Validate = validator.New()

func ValidateStruct(i interface{}, trans ut.Translator) (errs []response.InvalidField) {
	if err := Validate.Struct(i); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errs = translateErrorMessage(validationErrors, trans)
		}
	}

	return
}

func translateErrorMessage(validationErrors validator.ValidationErrors, trans ut.Translator) (errs []response.InvalidField) {
	var invalidField response.InvalidField
	for _, err := range validationErrors {
		invalidField = response.InvalidField{
			Field:  err.Field(),
			Reason: err.Translate(trans),
		}
		errs = append(errs, invalidField)
	}

	return
}
