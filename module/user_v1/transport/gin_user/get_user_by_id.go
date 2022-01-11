package gin_user

import (
	"amarbank/mockdata"
	service2 "amarbank/module/user_v1/service"
	"amarbank/module/user_v1/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
)

func GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userUUID uuid.UUID
		var err error
		if userUUID, err = uuid.Parse(c.Param("id")); err != nil {
			panic(err)
		}

		store, _ := storage.NewUSerFileStorage(mockdata.DefaultPathToUser, os.ModeAppend)
		service := service2.NewGetUserByIDService(store)

		user, err := service.GetUserByID(c.Request.Context(), userUUID)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": user})

	}
}
