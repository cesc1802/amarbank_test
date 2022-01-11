package dto

import (
	app_error "amarbank/pkg/apperror"
	"amarbank/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestValidate_CreateUserRequest(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("space-between", validation.StringContainWhiteSpace)
	validate.SetTagName("binding")

	t.Run("Test Case Name field must be have lastname and firstname", func(t *testing.T) {
		req := CreateUserRequest{
			Name:   "CescNguyen",
			Dob:    "18-02-1993",
			Gender: "male",
		}
		err := validate.Struct(req)
		if err != nil {

			if ve, ok := err.(validator.ValidationErrors); ok {

				appVE := app_error.HandleValidationErrors(ve)
				for _, e := range appVE.VE {
					assert.Equal(t, e.Field, "Name")
					assert.Equal(t, e.Tag, "space-between")
				}
				return
			}
		}
	})

	t.Run("test case create user request successfully", func(t *testing.T) {
		req := CreateUserRequest{
			Name:   "Cesc Nguyen",
			Dob:    "18-02-1993",
			Gender: "male",
		}
		err := validate.Struct(req)
		assert.Equal(t, err, nil)
	})

	t.Run("test case Gender field is not MALE or FEMALE", func(t *testing.T) {
		req := CreateUserRequest{
			Name:   "Cesc Nguyen",
			Dob:    "18-02-1993",
			Gender: "male-123",
		}
		err := validate.Struct(req)
		if err != nil {

			if ve, ok := err.(validator.ValidationErrors); ok {
				appVE := app_error.HandleValidationErrors(ve)
				for _, e := range appVE.VE {
					assert.Equal(t, e.Field, "Gender")
					assert.Equal(t, e.Tag, "oneof")
				}
				return
			}
		}
	})

}
