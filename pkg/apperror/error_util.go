package app_error

import (
	"github.com/go-playground/validator/v10"
)

func translateToAppVE(
	valErrors validator.ValidationErrors) []ValidationErrorField {

	res := make([]ValidationErrorField, len(valErrors))
	for i, valErr := range valErrors {
		res[i] = ValidationErrorField{
			Field: valErr.Field(),
			Tag:   valErr.Tag(),
		}
	}
	return res
}

func HandleValidationErrors(valErrors validator.ValidationErrors) *AppError {
	appErr := ValidationError(
		"ERR_VALIDATION_REQUEST",
		"ERR_VALIDATION_REQUEST",
		translateToAppVE(valErrors),
	)
	return appErr
}

func HandleAppError(err error) *AppError {
	appErr := err.(*AppError)
	return appErr

}

func MustError(err error) {
	if err != nil {
		panic(err)
	}
}
