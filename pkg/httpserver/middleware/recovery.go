package middleware

import (
	app_error "amarbank/pkg/apperror"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				if ve, ok := err.(validator.ValidationErrors); ok {
					appVE := app_error.HandleValidationErrors(ve)
					c.AbortWithStatusJSON(appVE.StatusCode, appVE)
					return
				}

				if appErr, ok := err.(*app_error.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, app_error.HandleAppError(appErr))
					return
				}

				appErr := app_error.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
				return
			}
		}()

		c.Next()
	}
}
