package gin_user

import (
	"amarbank/mockdata"
	"amarbank/module/user_v1/dto"
	service2 "amarbank/module/user_v1/service"
	"amarbank/module/user_v1/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.CreateUserRequest

		if err := c.ShouldBind(&input); err != nil {
			panic(err)
		}

		store, _ := storage.NewUSerFileStorage(mockdata.DefaultPathToUser, os.ModePerm)
		service := service2.NewCreateUserService(store)

		if err := service.CreateUser(c.Request.Context(), &input); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
