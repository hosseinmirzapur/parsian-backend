package validation

import (
	"github.com/go-playground/validator/v10"
)

func ValidateData(data interface{}) (map[string]string, bool) {
	v := validator.New()
	validationErrors := map[string]string{}

	err := v.Struct(data)

	if err != nil {
		for _, item := range err.(validator.ValidationErrors) {
			validationErrors[item.Field()] = item.Tag()
		}

		return validationErrors, false
	}

	return validationErrors, true

}
