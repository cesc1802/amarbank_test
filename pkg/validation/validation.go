package validation

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func StringContainWhiteSpace(field validator.FieldLevel) bool {
	return strings.Contains(strings.TrimSpace(field.Field().String()), " ")
}
